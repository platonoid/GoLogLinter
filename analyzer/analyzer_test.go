package analyzer_test

import (
	"testing"

	a "myloggerlint/analyzer"
	c "myloggerlint/config"
	r "myloggerlint/rules"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestRules(t *testing.T) {
	cfg, err := c.Load("../loggerlint.yaml")
	if err != nil {
		panic(err)
	}

	loggers, _ := c.BuildLoggers(cfg)

	if err := loggers.AddLogger("Logger", "testlog"); err != nil {
		t.Fatal(err)
	}

	if err := loggers.AddLoggerMethod("Logger", "Info"); err != nil {
		t.Fatal(err)
	}

	sensitive := r.NewSensitiveDataRule([]string{
		"password",
		"token",
	})

	capital := &r.CapitalLetterRule{}
	english := &r.EnglishOnlyRule{}
	symbols := &r.SpecialSymbolsRule{}

	sensitive.SetNext(capital)
	capital.SetNext(english)
	english.SetNext(symbols)

	analyzer := a.NewLoggerAnalyzer(loggers, sensitive)

	analysistest.Run(t, analysistest.TestData(), analyzer,
		"capital",
		"english",
		"specialsymbols",
		"sensetivedata",
	)
}
