// Package gomegax provides wrapper types over types from [gomega](https://pkg.go.dev/github.com/onsi/gomega) and [testing](https://pkg.go.dev/testing) packages.
package gomegax

import (
	"testing"

	"github.com/onsi/gomega"
)

// NewTB constructs a new TB.
func NewTB[T HierarchicalTest[T]](t T) TB[T] {
	return TB[T]{
		gomega.NewWithT(t),
		t,
		t,
	}
}

// ---

// TB combines gomega.WithT and testing.TB.
type TB[T HierarchicalTest[T]] struct {
	*gomega.WithT
	testing.TB
	inner T
}

// Inner returns the inner wrapped object.
func (t TB[T]) Inner() T {
	return t.inner
}

// Run runs a child test with the give name.
func (t TB[T]) Run(name string, f func(t TB[T])) bool {
	return t.inner.Run(name, func(t T) {
		f(NewTB(t))
	})
}

// Fail marks the function as having failed but continues execution.
func (t TB[T]) Fail() {
	t.TB.Fail()
}
