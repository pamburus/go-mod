package qb

func Value(val any) ValueExpression {
	return ValueExpression{val}
}

type ValueExpression struct {
	val any
}

func (v ValueExpression) BuildExpression(b Builder) error {
	b.AppendArg(v.val)

	return nil
}
