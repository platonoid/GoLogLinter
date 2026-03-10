package rules

import (
	"go/ast"
	"go/token"
	"strconv"
	"unicode"
)

type CapitalLetterRule struct {
	BaseRule
}

func (r *CapitalLetterRule) Check(args *CallArgs) {
	lit, ok := args.Arg.(*ast.BasicLit)
	if !ok {
		r.callNext(args)
		return
	}

	if lit.Kind != token.STRING {
		r.callNext(args)
		return
	}

	msg, err := strconv.Unquote(lit.Value)
	if err != nil || msg == "" {
		r.callNext(args)
		return
	}

	first := []rune(msg)[0]

	if !unicode.IsUpper(first) {
		args.Pass.Reportf(args.Arg.Pos(), "log message should start with uppercase letter")
	}

	r.callNext(args)
}
