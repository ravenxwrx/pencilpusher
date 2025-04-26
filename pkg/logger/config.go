package logger

const (
	LogTypeText = "text"
	LogTypeJSON = "json"

	LogLevelDebug = "debug"
	LogLevelInfo  = "info"
	LogLevelWarn  = "warn"
	LogLevelError = "error"
)

var (
	logLevel  string = LogLevelInfo
	logFormat string = LogTypeText
)

func LogLevel() string {
	return logLevel
}

func SetLogLevel(level string) {
	logLevel = level
}

func LogFormat() string {
	return logFormat
}

func SetLogFormat(format string) {
	logFormat = format
}
