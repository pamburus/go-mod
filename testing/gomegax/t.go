package gomegax

import (
	"testing"

	"github.com/onsi/gomega"
)

// NewT constructs a new T.
func NewT(t *testing.T) T {
	return T{
		gomega.NewWithT(t),
		t,
	}
}

// ---

// T combines gomega.WithT and testing.T.
type T struct {
	*gomega.WithT
	*testing.T
}

// Run runs a child test with the give name.
func (t T) Run(name string, f func(t T)) bool {
	return t.T.Run(name, func(t *testing.T) {
		f(NewT(t))
	})
}

// Fail marks the function as having failed but continues execution.
func (t T) Fail() {
	t.T.Fail()
}
