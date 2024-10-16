package qb

// Expression is an abstract SQL expression.
type Expression interface {
	BuildExpression(Builder, ExpressionOptions) error
}

type ExpressionOptions interface {
	AliasOptions
	sealedExpressionOptions()
}

// ---

func DefaultExpressionOptions() ExpressionOptions {
	return defaultExpressionOptionsInstance
}

// ---

var defaultExpressionOptionsInstance = &defaultExpressionOptions{}

// ---

type defaultExpressionOptions struct {
	defaultAliasOptions
}

func (*defaultExpressionOptions) sealedExpressionOptions() {}
