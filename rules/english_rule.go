package rules

import (
	"go/ast"
	"go/token"
	"strconv"
	"unicode"
)

type EnglishOnlyRule struct {
	BaseRule
}

func (r *EnglishOnlyRule) Check(args *CallArgs) {
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
		if unicode.IsLetter(ch) && !unicode.In(ch, unicode.Latin) {
			args.Pass.Reportf(args.Arg.Pos(), "log message should contain only English letters")
			break
		}
	}

	r.callNext(args)
}
