package logging

type RawLogger interface {
	Log(args ...any)
	Logf(format string, args ...any)
}

func DiscardRawLogger() RawLogger {
	return discardRawLoggerInstance
}

func RawPrefix(prefix string, logger RawLogger) RawLogger {
	return &rawPrefixLogger{prefix, logger}
}

// ---

var discardRawLoggerInstance RawLogger = discardRawLogger{}

type discardRawLogger struct{}

func (discardRawLogger) Log(args ...any)                 {}
func (discardRawLogger) Logf(format string, args ...any) {}

// ---

type rawPrefixLogger struct {
	prefix string
	logger RawLogger
}

func (l *rawPrefixLogger) Log(args ...any) {
	l.logger.Log(append([]any{l.prefix}, args...)...)
}

func (l *rawPrefixLogger) Logf(format string, args ...any) {
	l.logger.Logf(l.prefix+" "+format, args...)
}
