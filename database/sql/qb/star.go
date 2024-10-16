package qb

func Star() StarSymbol {
	return StarSymbol{}
}

// ---

type StarSymbol struct{}

func (StarSymbol) BuildExpression(b Builder, _ ExpressionOptions) error {
	b.AppendByte('*')

	return nil
}
