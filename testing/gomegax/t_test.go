package gomegax_test

import (
	"testing"

	"github.com/onsi/gomega"

	"github.com/pamburus/go-mod/testing/gomegax"
)

func TestT(tt *testing.T) {
	t := gomegax.NewT(tt)
	t.Expect(t.Failed()).To(gomega.BeFalse())
	t.Expect(1).To(gomega.Equal(1))

	t.Run("child", func(t gomegax.T) {
		t.Expect(t.Failed()).To(gomega.BeFalse())
		t.Expect(1).To(gomega.Equal(1))
	})
}
