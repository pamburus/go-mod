package qb

// Statement is an abstract SQL Statement.
type Statement interface {
	Build(Builder, StatementOptions) error
}

type StatementOptions interface {
	sealedStatementOptions()
}

// ---

func DefaultStatementOptions() StatementOptions {
	return defaultStatementOptionsInstance
}

// ---

var defaultStatementOptionsInstance = &defaultStatementOptions{}

// ---

type defaultStatementOptions struct{}

func (*defaultStatementOptions) sealedStatementOptions() {}
