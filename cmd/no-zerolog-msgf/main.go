package main

import (
	"github.com/ACK101101/nozerologmsgf/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.Analyzer}, nil
}
