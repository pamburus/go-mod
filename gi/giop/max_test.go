package giop_test

import (
	"math"
	"net/netip"
	"testing"

	"github.com/pamburus/go-mod/gi/giop"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, 3, giop.Max(2, 3))
		assert.Equal(t, 2, giop.Max(2, -3))
	})

	t.Run("Float", func(t *testing.T) {
		assert.InDelta(t, 2.5, giop.Max(2.5, 2.5), 1e-9)
		assert.InDelta(t, 2.5, giop.Max(2.5, -3.5), 1e-9)
		assert.InDelta(t, 3.5, giop.Max(2.5, 3.5), 1e-9)
	})
}

func TestMaxBy(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		mod2 := func(a int) int {
			return a % 2
		}

		assert.Equal(t, 3, giop.MaxBy(mod2)(2, 3))
		assert.Equal(t, 3, giop.MaxBy(mod2)(3, 4))
	})

	t.Run("Float", func(t *testing.T) {
		mod2 := func(a float64) float64 {
			return math.Mod(a, 2)
		}

		assert.InDelta(t, 3.0, giop.MaxBy(mod2)(2.0, 3.0), 1e-9)
		assert.InDelta(t, 3.0, giop.MaxBy(mod2)(3.0, 4.0), 1e-9)
	})
}

func TestMaxByLess(t *testing.T) {
	t.Run("Addr", func(t *testing.T) {
		addr1 := netip.MustParseAddr("192.168.1.1")
		addr2 := netip.MustParseAddr("192.168.1.2")

		assert.Equal(t, addr2, giop.MaxByLess(addr1, addr2))
		assert.Equal(t, addr2, giop.MaxByLess(addr2, addr1))
	})
}

func TestMaxByCompare(t *testing.T) {
	t.Run("Addr", func(t *testing.T) {
		addr1 := netip.MustParseAddr("192.168.1.1")
		addr2 := netip.MustParseAddr("192.168.1.2")

		assert.Equal(t, addr2, giop.MaxByCompare(addr1, addr2))
		assert.Equal(t, addr2, giop.MaxByCompare(addr2, addr1))
	})
}

func TestMaxByLessFunc(t *testing.T) {
	t.Run("StringLength", func(t *testing.T) {
		str1 := "short"
		str2 := "longer"

		less := func(a, b string) bool {
			return len(a) < len(b)
		}

		assert.Equal(t, str2, giop.MaxByLessFunc(less)(str1, str2))
		assert.Equal(t, str2, giop.MaxByLessFunc(less)(str2, str1))
	})
}

func TestMaxByCompareFunc(t *testing.T) {
	t.Run("StringLength", func(t *testing.T) {
		str1 := "short"
		str2 := "longer"

		compare := func(a, b string) int {
			return len(a) - len(b)
		}

		assert.Equal(t, str2, giop.MaxByCompareFunc(compare)(str1, str2))
		assert.Equal(t, str2, giop.MaxByCompareFunc(compare)(str2, str1))
	})
}
