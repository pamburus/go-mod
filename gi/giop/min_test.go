package giop_test

import (
	"cmp"
	"math"
	"net/netip"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/giop"
)

func TestMin(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, 2, giop.Min(2, 3))
		assert.Equal(t, -3, giop.Min(2, -3))
	})

	t.Run("Float", func(t *testing.T) {
		assert.InDelta(t, 2.5, giop.Min(2.5, 2.5), 1e-9)
		assert.InDelta(t, -3.5, giop.Min(2.5, -3.5), 1e-9)
	})
}

func TestMinBy(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		mod2 := func(a int) int {
			return a % 2
		}

		assert.Equal(t, 2, giop.MinBy(mod2)(2, 3))
		assert.Equal(t, 4, giop.MinBy(mod2)(3, 4))
	})

	t.Run("Float", func(t *testing.T) {
		mod2 := func(a float64) float64 {
			return math.Mod(a, 2)
		}

		assert.InDelta(t, 2.0, giop.MinBy(mod2)(2.0, 3.0), 1e-9)
		assert.InDelta(t, 4.0, giop.MinBy(mod2)(3.0, 4.0), 1e-9)
	})
}

func TestMinByLess(t *testing.T) {
	t.Run("Addr", func(t *testing.T) {
		addr1 := netip.MustParseAddr("192.168.1.1")
		addr2 := netip.MustParseAddr("192.168.1.2")

		assert.Equal(t, addr1, giop.MinByLess(addr1, addr2))
		assert.Equal(t, addr1, giop.MinByLess(addr2, addr1))
	})
}

func TestMinByLessFunc(t *testing.T) {
	t.Run("StringLength", func(t *testing.T) {
		less := func(a, b string) bool {
			return len(a) < len(b)
		}

		assert.Equal(t, "short", giop.MinByLessFunc(less)("short", "longer"))
		assert.Equal(t, "tiny", giop.MinByLessFunc(less)("tiny", "small"))
	})
}

func TestMinByCompare(t *testing.T) {
	t.Run("Addr", func(t *testing.T) {
		addr1 := netip.MustParseAddr("192.168.1.1")
		addr2 := netip.MustParseAddr("192.168.1.2")

		assert.Equal(t, addr1, giop.MinByCompare(addr1, addr2))
		assert.Equal(t, addr1, giop.MinByCompare(addr2, addr1))
	})
}

func TestMinByCompareFunc(t *testing.T) {
	t.Run("Float", func(t *testing.T) {
		assert.InDelta(t, 2.5, giop.MinByCompareFunc(cmp.Compare[float64])(2.5, 3.5), 1e-9)
		assert.InDelta(t, -3.5, giop.MinByCompareFunc(cmp.Compare[float64])(2.5, -3.5), 1e-9)
	})
}
