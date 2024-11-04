package backend

import (
	"context"
	"net/url"
)

type Backend interface {
	Start(context.Context, Options) (Server, StopFunc, error)
	BackendProvider
}

type Server interface {
	Context() context.Context
	URL() *url.URL
}

type StopFunc func(context.Context) error

type Options struct {
	Password string
	Port     uint16
}

type BackendProvider interface {
	Backend() Backend
}
