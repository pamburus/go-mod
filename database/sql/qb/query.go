package qb

// Query is an abstract SQL Query.
type Query interface {
	BuildQuery(Builder, QueryOptions) error
}

type QueryOptions interface {
	sealedQueryOptions()
}

// ---

func DefaultQueryOptions() QueryOptions {
	return defaultQueryOptionsInstance
}

// ---

var defaultQueryOptionsInstance = &defaultQueryOptions{}

// ---

type defaultQueryOptions struct{}

func (*defaultQueryOptions) sealedQueryOptions() {}
