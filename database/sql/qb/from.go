package qb

// FromItem is an abstract SQL from item.
type FromItem interface {
	BuildFromItem(Builder, FromItemOptions) error
}

type FromItemOptions interface {
	AliasOptions
	sealedFromItemOptions()
}

// ---

func DefaultFromItemOptions() FromItemOptions {
	return defaultFromItemOptionsInstance
}

// ---

var defaultFromItemOptionsInstance = &defaultFromItemOptions{}

// ---

type defaultFromItemOptions struct {
	defaultAliasOptions
}

func (*defaultFromItemOptions) sealedFromItemOptions() {}
