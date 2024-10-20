package qb

import (
	"errors"
	"strconv"
)

func Select(what ...Expression) SelectQuery {
	return SelectQuery{what: what}
}

// ---

type SelectQuery struct {
	what  []Expression
	from  []FromItem
	where Condition
	order order
	limit int
	alias string
}

func (s SelectQuery) From(items ...FromItem) SelectQuery {
	s.from = items

	return s
}

func (s SelectQuery) Where(condition Condition) SelectQuery {
	s.where = And(s.where, condition)

	return s
}

func (s SelectQuery) OrderBy(options ...OrderOption) SelectQuery {
	for _, option := range options {
		option(&s.order)
	}

	return s
}

func (s SelectQuery) Limit(n int) SelectQuery {
	s.limit = n

	return s
}

func (s SelectQuery) As(alias string) SelectQuery {
	s.alias = alias

	return s
}

func (s SelectQuery) BuildFromItem(b Builder, options FromItemOptions) error {
	build := func(b Builder) error {
		b.AppendByte('(')
		err := s.BuildQuery(b, DefaultQueryOptions())
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

func (s SelectQuery) BuildQuery(b Builder, _ QueryOptions) error {
	b.AppendString("SELECT")

	for i, item := range s.what {
		if i > 0 {
			b.AppendByte(',')
		}
		b.AppendByte(' ')

		err := item.BuildExpression(b, optSelectWhat)
		if err != nil {
			return err
		}
	}

	if len(s.from) != 0 {
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
		b.AppendString(strconv.Itoa(s.limit))
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
	_ Query    = SelectQuery{}
	_ FromItem = SelectQuery{}
)
