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
	alias string
}

func (s SelectStatement) From(items ...FromItem) SelectStatement {
	s.from = items

	return s
}

func (s SelectStatement) Where(condition Condition) SelectStatement {
	s.where = And(s.where, condition)

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

func (s SelectStatement) As(alias string) SelectStatement {
	s.alias = alias

	return s
}

func (s SelectStatement) BuildFromItem(b Builder, options FromItemOptions) error {
	build := func(b Builder) error {
		b.AppendByte('(')
		err := s.BuildStatement(b, DefaultStatementOptions())
		if err != nil {
			return err
		}
		b.AppendByte(')')

		return nil
	}

	if s.alias == "" && options.AliasApplicable() {
		return errors.New("alias is required")
	}

	return as{build, s.alias, options}.build(b)
}

func (s SelectStatement) BuildStatement(b Builder, _ StatementOptions) error {
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

			err := item.BuildExpression(b, optSelectWhat)
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

		err := item.BuildFromItem(b, optSelectFrom)
		if err != nil {
			return err
		}
	}

	if s.where != nil {
		b.AppendString(" WHERE ")
		err := s.where.BuildCondition(b, DefaultConditionOptions())
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

type selectWhatOptions struct {
	defaultExpressionOptions
}

func (*selectWhatOptions) AliasApplicable() bool {
	return true
}

// ---

type selectFromOptions struct {
	defaultFromItemOptions
}

func (*selectFromOptions) AliasApplicable() bool {
	return true
}

// ---

var (
	optSelectWhat = &selectWhatOptions{}
	optSelectFrom = &selectFromOptions{}
)

var (
	_ Statement = SelectStatement{}
	_ FromItem  = SelectStatement{}
)
