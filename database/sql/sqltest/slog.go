package sqltest

import (
	"context"
	"log/slog"
)

type discardLogHandler struct{}

func (discardLogHandler) Enabled(context.Context, slog.Level) bool {
	return false
}

func (discardLogHandler) Handle(context.Context, slog.Record) error {
	return nil
}

func (discardLogHandler) WithAttrs([]slog.Attr) slog.Handler {
	return discardLogHandler{}
}

func (discardLogHandler) WithGroup(string) slog.Handler {
	return discardLogHandler{}
}

// ---

var (
	_ slog.Handler = discardLogHandler{}
)
