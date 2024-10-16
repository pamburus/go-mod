package qb

// Builder is an interface for building SQL queries.
type Builder interface {
	AppendByte(byte)
	AppendString(string)
	AppendArg(any) error
	AppendRawArgs(...any) error
}

// BuilderResult is an interface for the result of building a SQL query.
type BuilderResult interface {
	Result() (string, []any)
}

type BuilderWithResult interface {
	Builder
	BuilderResult
}

// Build builds the SQL query.
func Build[F func() B, B BuilderWithResult](f F, q Statement) (string, []any, error) {
	b := f()

	err := q.BuildStatement(b, DefaultStatementOptions())
	if err != nil {
		return "", nil, err
	}

	sql, args := b.Result()

	return sql, args, nil
}
