package qb

import "database/sql"

func Arg(val any) ArgValue {
	return ArgValue{val: val}
}

func NamedArg(name string, val any) ArgValue {
	return ArgValue{val: sql.Named(name, val)}
}

// ---

type ArgValue struct {
	val   any
	alias string
}

func (v ArgValue) As(alias string) ArgValue {
	v.alias = alias

	return v
}

func (v ArgValue) BuildExpression(b Builder, options ExpressionOptions) error {
	build := func(b Builder) error {
		return b.AppendArg(v.val)
	}

	return as{build, v.alias, options}.build(b)
}

// ---

var _ Expression = ArgValue{}
