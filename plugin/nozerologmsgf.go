package main

import (
	linters "github.com/Stork-Oracle/nozerologmsgf"
	"golang.org/x/tools/go/analysis"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{linters.NoZeroLogMsgfAnalyzer}, nil
}
