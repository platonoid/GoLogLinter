package main

import (
	a "myloggerlint/analyzer"
	c "myloggerlint/config"
	r "myloggerlint/rules"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	cfg, err := c.Load("loggerlint.yaml")
	if err != nil {
		panic(err)
	}

	loggers, err := c.BuildLoggers(cfg)
	if err != nil {
		panic(err)
	}

	sensitive := r.NewSensitiveDataRule([]string{"password", "token"})
	capital := &r.CapitalLetterRule{}
	english := &r.EnglishOnlyRule{}
	symbols := &r.SpecialSymbolsRule{}

	sensitive.SetNext(capital)
	capital.SetNext(english)
	english.SetNext(symbols)

	analyzer := a.NewLoggerAnalyzer(loggers, sensitive)
	unitchecker.Main(analyzer)
}
