package gomegax

import (
	"testing"

	"github.com/onsi/gomega"
)

// NewB constructs a new B.
func NewB(b *testing.B) B {
	return B{
		gomega.NewWithT(b),
		b,
	}
}

// ---

// B combines gomega.WithT and testing.B.
type B struct {
	*gomega.WithT
	*testing.B
}

// Run runs a child test with the give name.
func (b B) Run(name string, f func(b B)) bool {
	return b.B.Run(name, func(b *testing.B) {
		f(NewB(b))
	})
}

// Fail marks the function as having failed but continues execution.
func (b B) Fail() {
	b.B.Fail()
}
