package qb

type AliasOptions interface {
	AliasApplicable() bool
	sealedAliasOptions()
}

func DefaultAliasOptions() AliasOptions {
	return defaultAliasOptionsInstance
}

// ---

type as struct {
	buildItem func(Builder) error
	alias     string
	options   AliasOptions
}

func (a as) build(b Builder) error {
	err := a.buildItem(b)
	if err != nil {
		return err
	}

	if a.alias != "" && a.options.AliasApplicable() {
		b.AppendString(" AS ")
		b.AppendString(a.alias)
	}

	return nil
}

// ---

var defaultAliasOptionsInstance = &defaultAliasOptions{}

// ---

type defaultAliasOptions struct{}

func (*defaultAliasOptions) AliasApplicable() bool {
	return false
}

func (*defaultAliasOptions) sealedAliasOptions() {}
