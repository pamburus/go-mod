package logging

import (
	"bytes"
	"context"
	"io"
	"log/slog"
)

func DiscardLogger() *slog.Logger {
	return discardLoggerInstance
}

func DiscardHandler() slog.Handler {
	return discardHandlerInstance
}

func RawToStructured(logger RawLogger, level slog.Level) *slog.Logger {
	return slog.New(slog.NewTextHandler(&loggerWriter{logger: logger}, &slog.HandlerOptions{
		AddSource: false,
		Level:     level,
		ReplaceAttr: func(groups []string, attr slog.Attr) slog.Attr {
			if len(groups) == 0 && attr.Key == slog.TimeKey {
				if attr.Value.Kind() == slog.KindTime {
					return slog.Attr{
						Key:   attr.Key,
						Value: slog.StringValue(attr.Value.Time().Format("2006-01-02 15:04:05.000")),
					}
				}
			}

			return attr
		},
	}))
}

// ---

type discardHandler struct{}

func (discardHandler) Enabled(context.Context, slog.Level) bool {
	return false
}

func (discardHandler) Handle(context.Context, slog.Record) error {
	return nil
}

func (discardHandler) WithAttrs([]slog.Attr) slog.Handler {
	return discardHandler{}
}

func (discardHandler) WithGroup(string) slog.Handler {
	return discardHandler{}
}

// ---

type loggerWriter struct {
	logger RawLogger
	buf    bytes.Buffer
}

func (w *loggerWriter) Write(p []byte) (int, error) {
	if w.buf.Len() == 0 && len(p) > 0 && p[len(p)-1] == '\n' {
		w.logger.Log(string(p[:len(p)-1]))
	} else {
		_, _ = w.buf.Write(p)
		for {
			index := bytes.IndexByte(w.buf.Bytes(), '\n')
			if index < 0 {
				break
			}
			w.logger.Log(string(w.buf.Next(index)))
		}
	}

	return len(p), nil
}

// ---

var (
	discardHandlerInstance slog.Handler = discardHandler{}
	discardLoggerInstance  *slog.Logger = slog.New(discardHandlerInstance)
	_                      io.Writer    = &loggerWriter{}
)
