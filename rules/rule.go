package rules

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

type CallArgs struct {
	Arg  ast.Expr
	Pass *analysis.Pass
}

type Rule interface {
	Check(*CallArgs)
	SetNext(Rule)
}
