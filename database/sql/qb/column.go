package qb

func Column(name string) ColumnRef {
	return ColumnRef{name: name}
}

// ---

type ColumnRef struct {
	name  string
	table TableRef
	alias string
}

func (c ColumnRef) Table(name string) ColumnRef {
	c.table = Table(name)

	return c
}

func (c ColumnRef) As(alias string) ColumnRef {
	c.alias = alias

	return c
}

func (c ColumnRef) BuildExpression(b Builder, options ExpressionOptions) error {
	build := func(b Builder) error {
		if c.table.name != "" {
			c.table.BuildFromItem(b, DefaultFromItemOptions())
			b.AppendByte('.')
		}

		b.AppendString(c.name)

		return nil
	}

	return as{build, c.alias, options}.build(b)
}

// ---

var _ Expression = ColumnRef{}
