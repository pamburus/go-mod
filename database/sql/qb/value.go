package qb

func Value(val any) ValueExpression {
	return ValueExpression{val: val}
}

// ---

type ValueExpression struct {
	val   any
	alias string
}

func (v ValueExpression) As(alias string) ValueExpression {
	v.alias = alias

	return v
}

func (v ValueExpression) BuildExpression(b Builder, options ExpressionOptions) error {
	build := func(b Builder) error {
		b.AppendArg(v.val)

		return nil
	}

	return as{build, v.alias, options}.build(b)
}

// ---

var _ Expression = ValueExpression{}
