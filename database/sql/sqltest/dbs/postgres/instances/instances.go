package instances

import (
	"context"
	"net/url"
)

type ManagerProvider interface {
	Manager() Manager
}

type Manager interface {
	Start(context.Context, Options) (Instance, StopFunc, error)
	ManagerProvider
}

type Instance interface {
	Context() context.Context
	URL() *url.URL
}

type StopFunc func(context.Context) error

type Options struct {
	Password string
	Port     uint16
}
