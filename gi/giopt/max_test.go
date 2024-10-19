package giopt_test

import (
	"net/netip"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/giopt"
	"github.com/pamburus/go-mod/optional/optval"
)

func TestMax(t *testing.T) {
	assert.Equal(t,
		optval.None[int](),
		giopt.Max(slices.Values([]int{})),
	)

	assert.Equal(t,
		optval.New(7, true),
		giopt.Max(slices.Values([]int{5, 3, 6, 7, 4})),
	)
}

func TestMaxBy(t *testing.T) {
	assert.Equal(t,
		optval.None[int](),
		giopt.MaxBy(slices.Values([]int{}), func(v int) int {
			return -v
		}),
	)

	assert.Equal(t,
		optval.New(3, true),
		giopt.MaxBy(slices.Values([]int{5, 3, 6, 7, 4}), func(v int) int {
			return -v
		}),
	)
}

func TestMaxByLess(t *testing.T) {
	assert.Equal(t,
		optval.None[netip.Addr](),
		giopt.MaxByLess(slices.Values([]netip.Addr{})),
	)

	assert.Equal(t,
		optval.New(netip.MustParseAddr("192.168.3.2"), true),
		giopt.MaxByLess(slices.Values([]netip.Addr{
			netip.MustParseAddr("192.168.2.1"),
			netip.MustParseAddr("192.168.1.1"),
			netip.MustParseAddr("192.168.3.2"),
		})),
	)
}

func TestMaxByLessFunc(t *testing.T) {
	assert.Equal(t,
		optval.None[netip.Addr](),
		giopt.MaxByLessFunc(slices.Values([]netip.Addr{}), netip.Addr.Less),
	)

	assert.Equal(t,
		optval.New(netip.MustParseAddr("192.168.3.2"), true),
		giopt.MaxByLessFunc(slices.Values([]netip.Addr{
			netip.MustParseAddr("192.168.2.1"),
			netip.MustParseAddr("192.168.1.1"),
			netip.MustParseAddr("192.168.3.2"),
		}), netip.Addr.Less),
	)
}

func TestMaxByCompare(t *testing.T) {
	assert.Equal(t,
		optval.None[netip.Addr](),
		giopt.MaxByCompare(slices.Values([]netip.Addr{})),
	)

	assert.Equal(t,
		optval.New(netip.MustParseAddr("192.168.3.2"), true),
		giopt.MaxByCompare(slices.Values([]netip.Addr{
			netip.MustParseAddr("192.168.2.1"),
			netip.MustParseAddr("192.168.1.1"),
			netip.MustParseAddr("192.168.3.2"),
		})),
	)
}

func TestMaxByCompareFunc(t *testing.T) {
	assert.Equal(t,
		optval.None[netip.Addr](),
		giopt.MaxByCompareFunc(slices.Values([]netip.Addr{}), netip.Addr.Compare),
	)

	assert.Equal(t,
		optval.New(netip.MustParseAddr("192.168.3.2"), true),
		giopt.MaxByCompareFunc(slices.Values([]netip.Addr{
			netip.MustParseAddr("192.168.2.1"),
			netip.MustParseAddr("192.168.1.1"),
			netip.MustParseAddr("192.168.3.2"),
		}), netip.Addr.Compare),
	)
}
