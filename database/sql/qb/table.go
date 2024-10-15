package qb

type Table struct {
	Name string
}

func (t Table) BuildFromItem(b Builder) error {
	b.AppendString(t.Name)

	return nil
}

// ---

var _ FromItem = Table{}
