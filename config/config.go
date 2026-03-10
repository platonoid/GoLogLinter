package config

type Config struct {
	Loggers map[string]LoggerConfig `yaml:"loggers"`
	Rules   RulesConfig             `yaml:"rules"`
}

type LoggerConfig struct {
	Package string   `yaml:"package"`
	Methods []string `yaml:"methods"`
}

type RulesConfig struct {
	SensitiveData SensitiveDataConfig `yaml:"sensitive_data"`
}

type SensitiveDataConfig struct {
	Patterns []string `yaml:"patterns"`
}
