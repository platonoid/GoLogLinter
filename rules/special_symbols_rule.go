package rules

import (
	"go/ast"
	"go/token"
	"strconv"
	"unicode"
)

type SpecialSymbolsRule struct {
	BaseRule
}

func (r *SpecialSymbolsRule) Check(args *CallArgs) {
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

	for _, ch := range msg {
		if !(unicode.IsLetter(ch) || unicode.IsDigit(ch) || unicode.IsSpace(ch)) {
			args.Pass.Reportf(args.Arg.Pos(), "log message must contain only letters and numbers")
			break
		}
	}

	r.callNext(args)
}
