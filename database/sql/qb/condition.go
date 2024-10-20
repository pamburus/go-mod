package qb

import (
	"slices"

	"github.com/pamburus/go-mod/gi"
)

// Condition is an abstract SQL condition.
type Condition interface {
	BuildCondition(Builder) error
}

// ---

func And(conditions ...Condition) Condition {
	conditions = cleanupConditions(conditions)

	switch len(conditions) {
	case 0:
		return nil
	case 1:
		return conditions[0]
	default:
		return and{conditions}
	}
}

func Or(conditions ...Condition) Condition {
	conditions = cleanupConditions(conditions)

	switch len(conditions) {
	case 0:
		return nil
	case 1:
		return conditions[0]
	default:
		return or{conditions}
	}
}

func Not(condition Condition) Condition {
	if condition == nil {
		return nil
	}

	return not{condition}
}

func Equal(left, right Expression) Condition {
	return binaryCondition{left, "=", right}
}

func NotEqual(left, right Expression) Condition {
	return binaryCondition{left, "<>", right}
}

func Less(left, right Expression) Condition {
	return binaryCondition{left, "<", right}
}

func LessOrEqual(left, right Expression) Condition {
	return binaryCondition{left, "<=", right}
}

func Greater(left, right Expression) Condition {
	return binaryCondition{left, ">", right}
}

func GreaterOrEqual(left, right Expression) Condition {
	return binaryCondition{left, ">=", right}
}

func Like(left, right Expression) Condition {
	return binaryCondition{left, "LIKE", right}
}

// ---

type and struct {
	conditions []Condition
}

func (a and) BuildCondition(b Builder) error {
	b.AppendByte('(')
	for i, condition := range a.conditions {
		if i > 0 {
			b.AppendString(" AND ")
		}

		err := condition.BuildCondition(b)
		if err != nil {
			return err
		}
	}
	b.AppendByte(')')

	return nil
}

// ---

type or struct {
	conditions []Condition
}

func (o or) BuildCondition(b Builder) error {
	b.AppendByte('(')
	for i, condition := range o.conditions {
		if i > 0 {
			b.AppendString(" OR ")
		}

		err := condition.BuildCondition(b)
		if err != nil {
			return err
		}
	}
	b.AppendByte(')')

	return nil
}

// ---

type not struct {
	condition Condition
}

func (n not) BuildCondition(b Builder) error {
	b.AppendString("NOT ")

	return n.condition.BuildCondition(b)
}

// ---

type binaryCondition struct {
	left  Expression
	op    string
	right Expression
}

func (bc binaryCondition) BuildCondition(b Builder) error {
	err := bc.left.BuildExpression(b)
	if err != nil {
		return err
	}

	b.AppendByte(' ')
	b.AppendString(bc.op)
	b.AppendByte(' ')

	err = bc.right.BuildExpression(b)
	if err != nil {
		return err
	}

	return nil
}

// ---

func cleanupConditions(conditions []Condition) []Condition {
	return slices.AppendSeq(
		make([]Condition, 0, len(conditions)),
		gi.Filter(slices.Values(conditions), gi.IsNotZero),
	)
}
