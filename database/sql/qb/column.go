package qb

func Column(name string) ColumnBuilder {
	return ColumnBuilder{name: name}
}

// ---

type ColumnBuilder struct {
	name  string
	table TableBuilder
	alias string
}

func (c ColumnBuilder) Table(name string) ColumnBuilder {
	c.table = Table(name)

	return c
}

func (c ColumnBuilder) As(alias string) ColumnBuilder {
	c.alias = alias

	return c
}

func (c ColumnBuilder) Build(b Builder, options ExpressionOptions) error {
	build := func(b Builder) error {
		if c.table.name != "" {
			c.table.BuildFromItem(b)
			b.AppendByte('.')
		}

		b.AppendString(c.name)

		return nil
	}

	return as{build, c.alias, options}.build(b)
}

// ---

var _ Expression = ColumnBuilder{}
