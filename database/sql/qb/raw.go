package qb

func Raw(sql string, args ...any) RawExpression {
	return RawExpression{sql: sql, args: args}
}

// ---

type RawExpression struct {
	sql   string
	args  []any
	alias string
}

func (r RawExpression) As(alias string) RawExpression {
	r.alias = alias

	return r
}

func (r RawExpression) BuildExpression(b Builder, options ExpressionOptions) error {
	return r.build(b, options)
}

func (r RawExpression) BuildCondition(b Builder, options ConditionOptions) error {
	return r.build(b, DefaultAliasOptions())
}

func (r RawExpression) BuildFromItem(b Builder, options FromItemOptions) error {
	return r.build(b, options)
}

func (r RawExpression) BuildQuery(b Builder, options QueryOptions) error {
	return r.build(b, DefaultAliasOptions())
}

func (r RawExpression) build(b Builder, options AliasOptions) error {
	build := func(b Builder) error {
		return b.AppendRawExpr(r.sql, r.args...)
	}

	return as{build, r.alias, options}.build(b)
}

// ---

var (
	_ Expression = RawExpression{}
	_ Condition  = RawExpression{}
	_ FromItem   = RawExpression{}
	_ Query      = RawExpression{}
)
