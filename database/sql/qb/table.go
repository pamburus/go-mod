package qb

func Table(name string) TableRef {
	return TableRef{name: name}
}

// ---

type TableRef struct {
	name  string
	alias string
}

func (t TableRef) Column(name string) ColumnRef {
	return ColumnRef{name: name, table: t}
}

func (t TableRef) AllColumns() ColumnRef {
	return AllColumns().Table(t.name)
}

func (t TableRef) As(alias string) TableRef {
	t.alias = alias

	return t
}

func (t TableRef) BuildFromItem(b Builder, options FromItemOptions) error {
	build := func(b Builder) error {
		b.AppendString(t.name)

		return nil
	}

	return as{build, t.alias, options}.build(b)
}

// ---

var _ FromItem = TableRef{}
