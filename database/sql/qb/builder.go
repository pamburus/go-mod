package qb

// Builder is an interface for building SQL queries.
type Builder interface {
	AppendByte(byte)
	AppendString(string)
	AppendArg(any) error
	AppendRawArgs(...any) error
}
