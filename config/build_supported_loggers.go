package config

import l "myloggerlint/internal"

func BuildLoggers(cfg *Config) (*l.SupportedLoggers, error) {
	loggers := l.NewSupportedLoggers()

	for logger, info := range cfg.Loggers {

		if err := loggers.AddLogger(logger, info.Package); err != nil {
			return nil, err
		}

		for _, method := range info.Methods {
			if err := loggers.AddLoggerMethod(logger, method); err != nil {
				return nil, err
			}
		}
	}

	return loggers, nil
}
