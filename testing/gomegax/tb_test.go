package gomegax_test

import (
	"testing"

	"github.com/onsi/gomega"

	"github.com/pamburus/go-mod/testing/gomegax"
)

func TestTB(tt *testing.T) {
	tbi := tb{tt, false}
	t := gomegax.NewTB(&tbi)
	t.Expect(t.Failed()).To(gomega.BeFalse())
	t.Expect(1).To(gomega.Equal(1))

	t.Run("child", func(t gomegax.TB[*tb]) {
		t.Expect(t.Failed()).To(gomega.BeFalse())
		t.Expect(1).To(gomega.Equal(1))
	})

	t.Expect(t.Inner()).ToNot(gomega.BeZero())
	t.Expect(t.Failed()).To(gomega.BeFalse())
	t.Expect(tbi.failed).To(gomega.BeFalse())

	t.Fail()
	t.Expect(t.Failed()).To(gomega.BeFalse())
	t.Expect(tbi.failed).To(gomega.BeTrue())
}

// ---

type tb struct {
	*testing.T
	failed bool
}

func (t *tb) Fail() {
	t.failed = true
}

func (t *tb) Run(name string, f func(t *tb)) bool {
	return t.T.Run(name, func(tt *testing.T) {
		f(&tb{tt, false})
	})
}
