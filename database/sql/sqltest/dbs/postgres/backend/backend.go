package backend

import (
	"context"
	"log/slog"
	"net/url"
)

type Backend interface {
	Start(context.Context, Options) (Server, StopFunc, error)
}

type Server interface {
	URL() *url.URL
}

type StopFunc func(context.Context) error

type Options struct {
	Password string
	Port     uint16
	Logger   *slog.Logger
}
