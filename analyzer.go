package linter

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

var ErrorMsg = "Do not use zerolog .Msgf after zerolog .Error; include extra info in Event fields"

var NoZeroLogMsgfAnalyzer = &analysis.Analyzer{
	Name: "nozerologmsgf",
	Doc:  "Finds .Msgf use on a zerolog Event after chained .Error use",
	Run:  msgfLintRun,
}

// Run function in the MsgfLintAnalyzer implementation.
func msgfLintRun(pass *analysis.Pass) (interface{}, error) {
	var zerologEventType types.Type

	// Find the zerolog import and zerolog Event type
	for _, pkg := range pass.Pkg.Imports() {
		if pkg.Name() == "zerolog" {
			if eventType := pkg.Scope().Lookup("Event"); eventType != nil {
				zerologEventType = eventType.Type()
			}
		}
	}

	// If zerolog is not imported, lint does not apply
	if zerologEventType == nil {
		return nil, nil
	}

	// Apply lint rule to each node in each file
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			if isMsgf, selExpr := isMsgfExprNode(pass, node, zerologEventType); isMsgf {
				// Walk up the chain of method calls to see if it came from Error()
				if isErrorExprInChain(selExpr.X) {
					pass.Report(analysis.Diagnostic{
						Pos:     node.Pos(),
						End:     node.End(),
						Message: ErrorMsg,
					})
				}
			}

			return true
		})
	}

	return nil, nil
}

// See if node is the .Msgf selector expression on a zerolog.Event.
func isMsgfExprNode(pass *analysis.Pass, node ast.Node, underlyingReceiverType types.Type) (bool, *ast.SelectorExpr) {
	// Look for method calls
	callExpr, ok := node.(*ast.CallExpr)
	if !ok {
		return false, nil
	}

	// Check if it's a selector expression (x.y())
	selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false, nil
	}

	// Get the type information of the receiver (x in x.y())
	recvType := pass.TypesInfo.TypeOf(selExpr.X)
	if recvType == nil {
		return false, nil
	}

	// Get the underlying type, handling pointer types
	underType := recvType
	if ptr, ok := recvType.(*types.Pointer); ok {
		underType = ptr.Elem()
	}

	// Check if this is a zerolog.Event
	if !types.Identical(underType, underlyingReceiverType) {
		return false, nil
	}

	if selExpr.Sel.Name != "Msgf" {
		return false, nil
	}

	return true, selExpr
}

// Checks if the expression chain includes an Error() call.
func isErrorExprInChain(expr ast.Expr) bool {
	callExpr, ok := expr.(*ast.CallExpr)
	if !ok {
		return false
	}

	selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	if selExpr.Sel.Name == "Error" {
		return true
	}

	// Else, recur to end of chain
	return isErrorExprInChain(selExpr.X)
}
