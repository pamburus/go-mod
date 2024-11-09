package logging

type RawLogger interface {
	Log(args ...any)
	Logf(format string, args ...any)
}

func DiscardRawLogger() RawLogger {
	return discardRawLoggerInstance
}

func RawPrefix(prefix string, logger RawLogger) RawLogger {
	return &rawPrefixLogger{prefix, logger, help(logger)}
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
	helper
}

func (l *rawPrefixLogger) Log(args ...any) {
	help(l.logger).Helper()
	l.logger.Log(append([]any{l.prefix}, args...)...)
}

func (l *rawPrefixLogger) Logf(format string, args ...any) {
	help(l.logger).Helper()
	l.logger.Logf(l.prefix+" "+format, args...)
}

// ---

func help(logger RawLogger) helper {
	if h, ok := logger.(helper); ok {
		return h
	}

	logger.Logf("warning: logger %T does not implement helper", logger)

	return noHelper{}
}

// ---

type helper interface {
	Helper()
}

// ---

type noHelper struct{}

func (noHelper) Helper() {}
