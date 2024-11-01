package sqltest

import "context"

type Option[T Base[T]] func(*jointOptions[T])

func WithContext[T Base[T]](ctx func(T) context.Context) Option[T] {
	return func(o *jointOptions[T]) {
		o.ctx = ctx
	}
}

// ---

func joinOptions[T Base[T]](options []Option[T]) jointOptions[T] {
	o := jointOptions[T]{}
	for _, opt := range options {
		opt(&o)
	}

	return o
}

type jointOptions[T Base[T]] struct {
	ctx func(T) context.Context
}
