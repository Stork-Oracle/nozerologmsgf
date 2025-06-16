package nozerologmsgf

import (
	"go/ast"
	"go/types"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

var ErrorMsg = "Do not use zerolog .Msgf after zerolog .Error; include extra info in Event fields"

func init() {
	register.Plugin("nozerologmsgf", New)
}

type NoZerologMsgfPlugin struct{}

func New(settings any) (register.LinterPlugin, error) {
	return &NoZerologMsgfPlugin{}, nil
}

func (n *NoZerologMsgfPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		{
			Name: "nozerologmsgf",
			Doc:  "Finds .Msgf use on a zerolog Event after chained .Error use",
			Run:  n.msgfLintRun,
		},
	}, nil
}

func (n *NoZerologMsgfPlugin) GetLoadMode() string {
	return register.LoadModeSyntax
}

// Run function in the MsgfLintAnalyzer implementation.
func (n *NoZerologMsgfPlugin) msgfLintRun(pass *analysis.Pass) (any, error) {
	var zerologEventType types.Type

	// Safely check if package exists
	if pass.Pkg == nil {
		return nil, nil
	}

	// Find the zerolog import and zerolog Event type
	for _, pkg := range pass.Pkg.Imports() {
		if pkg == nil {
			continue
		}
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
			if isMsgf, selExpr := n.isMsgfExprNode(pass, node, zerologEventType); isMsgf {
				// Walk up the chain of method calls to see if it came from Error()
				if n.isErrorExprInChain(selExpr.X) {
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
func (n *NoZerologMsgfPlugin) isMsgfExprNode(pass *analysis.Pass, node ast.Node, underlyingReceiverType types.Type) (bool, *ast.SelectorExpr) {
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
func (n *NoZerologMsgfPlugin) isErrorExprInChain(expr ast.Expr) bool {
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
	return n.isErrorExprInChain(selExpr.X)
}
