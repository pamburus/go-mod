package qb

import (
	"errors"
	"strconv"
)

func Select(what ...Expression) SelectCommand {
	return SelectCommand{what: what}
}

// ---

type SelectCommand struct {
	what  []Expression
	from  []FromItem
	where Condition
	order order
	limit int
	alias string
}

func (s SelectCommand) From(items ...FromItem) SelectCommand {
	s.from = items

	return s
}

func (s SelectCommand) Where(condition Condition) SelectCommand {
	s.where = And(s.where, condition)

	return s
}

func (s SelectCommand) OrderBy(options ...OrderOption) SelectCommand {
	for _, option := range options {
		option(&s.order)
	}

	return s
}

func (s SelectCommand) Limit(n int) SelectCommand {
	s.limit = n

	return s
}

func (s SelectCommand) As(alias string) SelectCommand {
	s.alias = alias

	return s
}

func (s SelectCommand) BuildFromItem(b Builder, options FromItemOptions) error {
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

func (s SelectCommand) BuildQuery(b Builder, _ QueryOptions) error {
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
	_ Query    = SelectCommand{}
	_ FromItem = SelectCommand{}
)
