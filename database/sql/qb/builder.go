package qb

// Builder is an interface for building SQL queries.
type Builder interface {
	// AppendByt appends a byte to the SQL query.
	AppendByte(byte)

	// AppendString appends a string to the SQL query.
	AppendString(string)

	// AppendArg appends an argument to the SQL query.
	// The argument can be a sql.NamedArg, pgx.NamedArgs, pgx.StrictNamedArgs, or any other type.
	AppendArg(any) error

	// AppendRawExpr appends a raw expression to the SQL query that can contain placeholders and arguments.
	// The arguments can be a sql.NamedArg, pgx.NamedArgs, pgx.StrictNamedArgs, or any other type.
	// The placeholders are replaced with the arguments in the order they are provided.
	// The format of the placeholders is implementation dependent, so using arguments in this method is not portable.
	AppendRawExpr(expr string, args ...any) error

	sealed()
}

// BuilderStub returns a stub implementation of Builder.
func BuilderStub() Builder {
	return builderStubInstance
}

// ---

var builderStubInstance = &builderStub{}

// ---

type builderStub struct{}

func (builderStub) AppendByte(byte) {
	panic(ErrNotImplemented)
}

func (builderStub) AppendString(string) {
	panic(ErrNotImplemented)
}

func (builderStub) AppendArg(any) error {
	return ErrNotImplemented
}

func (builderStub) AppendRawExpr(string, ...any) error {
	return ErrNotImplemented
}

func (builderStub) sealed() {}

// ---

var _ Builder = builderStub{}
