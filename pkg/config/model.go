package config

const (
	LogTypeText = "text"
	LogTypeJSON = "json"

	LogLevelDebug = "debug"
	LogLevelInfo  = "info"
	LogLevelWarn  = "warn"
	LogLevelError = "error"
)

type Config struct {
	Logging Logging `yaml:"logging"`
}

type Logging struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}
