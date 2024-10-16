package qb

func Table(name string) TableBuilder {
	return TableBuilder{name: name}
}

// ---

type TableBuilder struct {
	name string
}

func (t TableBuilder) Column(name string) ColumnBuilder {
	return ColumnBuilder{name: name, table: t}
}

func (t TableBuilder) BuildFromItem(b Builder) error {
	b.AppendString(t.name)

	return nil
}

// ---

var _ FromItem = TableBuilder{}
