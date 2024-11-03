package logctx

import (
	"context"
	"log/slog"

	"github.com/pamburus/go-mod/database/sql/sqltest/util/logging"
)

func New(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxKeyLogger, logger)
}

func Get(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(ctxKeyLogger).(*slog.Logger); ok {
		return logger
	}

	return logging.DiscardLogger()
}

// ---

type ctxKey int

const ctxKeyLogger ctxKey = iota
