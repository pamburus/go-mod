package qb

func Asc(expr Expression) OrderOption {
	return func(o *order) {
		o.items = append(o.items, orderItem{expr: expr})
	}
}

func Desc(expr Expression) OrderOption {
	return func(o *order) {
		o.items = append(o.items, orderItem{expr: expr, desc: true})
	}
}

func NullsFirst() OrderOption {
	return func(o *order) {
		o.nullsFirst = true
		o.nullsLast = false
	}
}

func NullsLast() OrderOption {
	return func(o *order) {
		o.nullsLast = true
		o.nullsFirst = false
	}
}

// ---

type OrderOption func(*order)

// ---

type order struct {
	items      []orderItem
	nullsFirst bool
	nullsLast  bool
}

func (o order) Build(b Builder) error {
	if len(o.items) == 0 {
		return nil
	}

	b.AppendString(" ORDER BY ")

	for i, item := range o.items {
		if i > 0 {
			b.AppendString(", ")
		}

		err := item.expr.Build(b, DefaultExpressionOptions())
		if err != nil {
			return err
		}

		if item.desc {
			b.AppendString(" DESC")
		}
	}

	if o.nullsFirst {
		b.AppendString(" NULLS FIRST")
	}

	if o.nullsLast {
		b.AppendString(" NULLS LAST")
	}

	return nil
}

// ---

type orderItem struct {
	expr Expression
	desc bool
}
