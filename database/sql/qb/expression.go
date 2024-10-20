package qb

// Expression is an abstract SQL expression.
type Expression interface {
	Build(Builder, ExpressionOptions) error
}

type ExpressionOptions interface {
	AliasOptions
	sealed()
}

// ---

func DefaultExpressionOptions() ExpressionOptions {
	return defaultExpressionOptionsInstance
}

// ---

var defaultExpressionOptionsInstance = &defaultExpressionOptions{}

type defaultExpressionOptions struct{}

func (*defaultExpressionOptions) AliasApplicable() bool {
	return false
}

func (*defaultExpressionOptions) sealed() {}

// ---

type aliasedExpression as

func (a aliasedExpression) BuildExpression(b Builder) error {
	return as(a).build(b)
}
