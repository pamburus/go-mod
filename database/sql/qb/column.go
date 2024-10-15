package qb

type Column struct {
	Name  string
	Table string
}

func (c Column) BuildExpression(b Builder) error {
	if c.Table != "" {
		b.AppendString(c.Table)
		b.AppendByte('.')
	}

	b.AppendString(c.Name)

	return nil
}

// ---

var _ Expression = Column{}
