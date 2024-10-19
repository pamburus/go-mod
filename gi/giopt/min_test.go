package giopt_test

import (
	"net/netip"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/giopt"
	"github.com/pamburus/go-mod/optional/optval"
)

func TestMin(t *testing.T) {
	assert.Equal(t,
		optval.None[int](),
		giopt.Min(slices.Values([]int{})),
	)

	assert.Equal(t,
		optval.New(3, true),
		giopt.Min(slices.Values([]int{5, 3, 6, 7, 4})),
	)
}

func TestMinBy(t *testing.T) {
	assert.Equal(t,
		optval.None[int](),
		giopt.MinBy(slices.Values([]int{}), func(v int) int {
			return -v
		}),
	)

	assert.Equal(t,
		optval.New(7, true),
		giopt.MinBy(slices.Values([]int{5, 3, 6, 7, 4}), func(v int) int {
			return -v
		}),
	)
}

func TestMinByLess(t *testing.T) {
	assert.Equal(t,
		optval.None[netip.Addr](),
		giopt.MinByLess(slices.Values([]netip.Addr{})),
	)

	assert.Equal(t,
		optval.New(netip.MustParseAddr("192.168.1.1"), true),
		giopt.MinByLess(slices.Values([]netip.Addr{
			netip.MustParseAddr("192.168.2.1"),
			netip.MustParseAddr("192.168.1.1"),
			netip.MustParseAddr("192.168.3.2"),
		})),
	)
}

func TestMinByLessFunc(t *testing.T) {
	assert.Equal(t,
		optval.None[netip.Addr](),
		giopt.MinByLessFunc(slices.Values([]netip.Addr{}), netip.Addr.Less),
	)

	assert.Equal(t,
		optval.New(netip.MustParseAddr("192.168.1.1"), true),
		giopt.MinByLessFunc(slices.Values([]netip.Addr{
			netip.MustParseAddr("192.168.2.1"),
			netip.MustParseAddr("192.168.1.1"),
			netip.MustParseAddr("192.168.3.2"),
		}), netip.Addr.Less),
	)
}

func TestMinByCompare(t *testing.T) {
	assert.Equal(t,
		optval.None[netip.Addr](),
		giopt.MinByCompare(slices.Values([]netip.Addr{})),
	)

	assert.Equal(t,
		optval.New(netip.MustParseAddr("192.168.1.1"), true),
		giopt.MinByCompare(slices.Values([]netip.Addr{
			netip.MustParseAddr("192.168.2.1"),
			netip.MustParseAddr("192.168.1.1"),
			netip.MustParseAddr("192.168.3.2"),
		})),
	)
}

func TestMinByCompareFunc(t *testing.T) {
	assert.Equal(t,
		optval.None[netip.Addr](),
		giopt.MinByCompareFunc(slices.Values([]netip.Addr{}), netip.Addr.Compare),
	)

	assert.Equal(t,
		optval.New(netip.MustParseAddr("192.168.1.1"), true),
		giopt.MinByCompareFunc(slices.Values([]netip.Addr{
			netip.MustParseAddr("192.168.2.1"),
			netip.MustParseAddr("192.168.1.1"),
			netip.MustParseAddr("192.168.3.2"),
		}), netip.Addr.Compare),
	)
}
