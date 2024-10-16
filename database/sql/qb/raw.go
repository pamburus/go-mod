package qb

func Raw(sql string, args ...any) RawExpression {
	return RawExpression{sql, args}
}

// ---

type RawExpression struct {
	sql  string
	args []any
}

func (r RawExpression) BuildExpression(b Builder, _ ExpressionOptions) error {
	return r.build(b)
}

func (r RawExpression) BuildCondition(b Builder, _ ConditionOptions) error {
	return r.build(b)
}

func (r RawExpression) BuildFromItem(b Builder, _ FromItemOptions) error {
	return r.build(b)
}

func (r RawExpression) BuildStatement(b Builder, _ StatementOptions) error {
	return r.build(b)
}

func (r RawExpression) build(b Builder) error {
	b.AppendString(r.sql)
	for _, arg := range r.args {
		b.AppendArg(arg)
	}

	return nil
}

// ---

var (
	_ Expression = RawExpression{}
	_ Condition  = RawExpression{}
	_ FromItem   = RawExpression{}
	_ Statement  = RawExpression{}
)
