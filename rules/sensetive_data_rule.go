package rules

import (
	"go/ast"
	"regexp"
)

type SensitiveDataRule struct {
	BaseRule
	Patterns []*regexp.Regexp
}

func NewSensitiveDataRule(patterns []string) *SensitiveDataRule {
	re := make([]*regexp.Regexp, 0, len(patterns))
	for _, p := range patterns {
		re = append(re, regexp.MustCompile(p))
	}
	return &SensitiveDataRule{
		Patterns: re,
	}
}

func (r *SensitiveDataRule) checkIdent(ident *ast.Ident, args *CallArgs) {
	for _, pattern := range r.Patterns {
		if pattern.MatchString(ident.Name) {
			args.Pass.Reportf(ident.Pos(), "potentially sensitive variable '%s' detected", ident.Name)
			break
		}
	}
}

func (r *SensitiveDataRule) findIdentifiers(expr ast.Expr) []*ast.Ident {
	var idents []*ast.Ident

	switch v := expr.(type) {
	case *ast.Ident:
		idents = append(idents, v)
	case *ast.BinaryExpr:
		idents = append(idents, r.findIdentifiers(v.X)...)
		idents = append(idents, r.findIdentifiers(v.Y)...)
	case *ast.UnaryExpr:
		idents = append(idents, r.findIdentifiers(v.X)...)
	case *ast.ParenExpr:
		idents = append(idents, r.findIdentifiers(v.X)...)
	case *ast.CallExpr:
		for _, arg := range v.Args {
			idents = append(idents, r.findIdentifiers(arg)...)
		}
	case *ast.SelectorExpr:
		idents = append(idents, r.findIdentifiers(v.X)...)
	case *ast.StarExpr:
		idents = append(idents, r.findIdentifiers(v.X)...)
	case *ast.IndexExpr:
		idents = append(idents, r.findIdentifiers(v.X)...)
		idents = append(idents, r.findIdentifiers(v.Index)...)
	case *ast.SliceExpr:
		idents = append(idents, r.findIdentifiers(v.X)...)

		if v.Low != nil {
			idents = append(idents, r.findIdentifiers(v.Low)...)
		}
		if v.High != nil {
			idents = append(idents, r.findIdentifiers(v.High)...)
		}
		if v.Max != nil {
			idents = append(idents, r.findIdentifiers(v.Max)...)
		}
	case *ast.TypeAssertExpr:
		idents = append(idents, r.findIdentifiers(v.X)...)
	case *ast.CompositeLit:
		for _, elt := range v.Elts {
			idents = append(idents, r.findIdentifiers(elt)...)
		}
	case *ast.KeyValueExpr:
		idents = append(idents, r.findIdentifiers(v.Key)...)
		idents = append(idents, r.findIdentifiers(v.Value)...)
	}

	return idents
}

func (r *SensitiveDataRule) Check(args *CallArgs) {
	idents := r.findIdentifiers(args.Arg)

	for _, ident := range idents {
		r.checkIdent(ident, args)
	}

	r.callNext(args)
}
