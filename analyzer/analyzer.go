package analyzer

import (
	"go/ast"
	"go/types"
	l "myloggerlint/internal"
	r "myloggerlint/rules"

	"golang.org/x/tools/go/analysis"
)

func NewLoggerAnalyzer(loggers *l.SupportedLoggers, rules r.Rule) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "loglint",
		Doc:  "checks log messages for following the rules",
		Run: func(pass *analysis.Pass) (interface{}, error) {
			ctx := AnalyzerContext{
				pass:    pass,
				loggers: loggers,
				rules:   rules,
			}
			return run(ctx)
		},
	}
}

func run(context AnalyzerContext) (interface{}, error) {
	for _, file := range context.pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			methodName := sel.Sel.Name

			typ := context.pass.TypesInfo.TypeOf(sel.X)
			if typ == nil {
				return true
			}

			if ptr, ok := typ.(*types.Pointer); ok { //for zap
				typ = ptr.Elem()
			}

			named, ok := typ.(*types.Named)
			if !ok {
				return true
			}

			obj := named.Obj()
			if obj == nil || obj.Pkg() == nil {
				return true
			}

			loggerName := obj.Name()
			pkgPath := obj.Pkg().Path()

			if context.loggers.IsLoggerCall(loggerName, pkgPath, methodName) {
				if len(call.Args) > 0 {
					context.rules.Check(&r.CallArgs{
						Pass: context.pass,
						Arg:  call.Args[0],
					})
				}

			}

			return true
		})
	}
	return nil, nil
}
