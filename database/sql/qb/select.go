package qb

import "errors"

func Select(what ...Expression) SelectWhat {
	return SelectWhat{what}
}

// ---

type SelectWhat struct {
	what []Expression
}

func (s SelectWhat) From(items ...FromItem) SelectStatement {
	return SelectStatement{what: s.what, from: items}
}

// ---

type SelectStatement struct {
	what  []Expression
	from  []FromItem
	where Condition
	order order
	limit int
}

func (s SelectStatement) From(items ...FromItem) SelectStatement {
	s.from = items

	return s
}

func (s SelectStatement) Where(condition Condition) SelectStatement {
	s.where = condition

	return s
}

func (s SelectStatement) OrderBy(options ...OrderOption) SelectStatement {
	for _, option := range options {
		option(&s.order)
	}

	return s
}

func (s SelectStatement) Limit(n int) SelectStatement {
	s.limit = n

	return s
}

func (s SelectStatement) BuildStatement(b Builder) error {
	if len(s.from) == 0 {
		return errors.New("no FROM items")
	}

	b.AppendString("SELECT ")

	if len(s.what) == 0 {
		b.AppendString("*")
	} else {
		for i, item := range s.what {
			if i > 0 {
				b.AppendString(", ")
			}

			err := item.Build(b, selectWhatOptionsInstance)
			if err != nil {
				return err
			}
		}
	}

	b.AppendString(" FROM ")

	for i, item := range s.from {
		if i > 0 {
			b.AppendString(", ")
		}

		err := item.BuildFromItem(b)
		if err != nil {
			return err
		}
	}

	if s.where != nil {
		b.AppendString(" WHERE ")
		err := s.where.BuildCondition(b)
		if err != nil {
			return err
		}
	}

	if len(s.order.items) != 0 {
		err := s.order.Build(b)
		if err != nil {
			return err
		}
	}

	if s.limit != 0 {
		b.AppendString(" LIMIT ")
		b.AppendArg(s.limit)
	}

	return nil
}

// ---

var selectWhatOptionsInstance = &selectWhatOptions{}

// ---

type selectWhatOptions struct {
	defaultExpressionOptions
}

func (*selectWhatOptions) AliasApplicable() bool {
	return true
}
