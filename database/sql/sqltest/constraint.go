package sqltest

import "testing"

// ---

type Base[T testing.TB] interface {
	testing.TB
	Run(name string, f func(T)) bool
}
