package gomegax_test

import (
	"testing"

	"github.com/onsi/gomega"

	"github.com/pamburus/go-mod/testing/gomegax"
)

func TestB(t *testing.T) {
	g := gomega.NewWithT(t)
	b := gomegax.NewB(&testing.B{})
	b.Fail()
	g.Expect(b.Failed()).To(gomega.BeTrue())
	b.Expect(1).To(gomega.Equal(1))
}
