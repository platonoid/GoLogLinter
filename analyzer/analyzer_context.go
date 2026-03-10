package analyzer

import (
	l "myloggerlint/internal"
	r "myloggerlint/rules"

	"golang.org/x/tools/go/analysis"
)

type AnalyzerContext struct {
	pass    *analysis.Pass
	loggers *l.SupportedLoggers
	rules   r.Rule
}
