package config

import r "myloggerlint/rules"

func BuildSensetiveDataRule(cfg *Config) r.Rule {
	sensitive := r.NewSensitiveDataRule(cfg.Rules.SensitiveData.Patterns)

	return sensitive
}
