package optpair_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/optional/optpair"
	"github.com/pamburus/go-mod/optional/optval"
)

func TestNew(t *testing.T) {
	pair := optpair.New(1, "one", true)
	v1, v2, ok := pair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 1, v1)
	assert.Equal(t, "one", v2)

	pair = optpair.New(1, "one", false)
	assert.False(t, pair.IsSome())
}

func TestSome(t *testing.T) {
	pair := optpair.Some(1, "one")
	v1, v2, ok := pair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 1, v1)
	assert.Equal(t, "one", v2)
}

func TestNone(t *testing.T) {
	pair := optpair.None[int, string]()
	assert.False(t, pair.IsSome())
}

func TestFromValue(t *testing.T) {
	split := func(v int) (int, bool) {
		return v, v%2 != 0
	}

	value := optval.Some(1)
	pair := optpair.FromValue(value, split)
	v1, v2, ok := pair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 1, v1)
	assert.True(t, v2)

	value = optval.Some(2)
	pair = optpair.FromValue(value, split)
	v1, v2, ok = pair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 2, v1)
	assert.False(t, v2)

	value = optval.None[int]()
	pair = optpair.FromValue(value, split)
	assert.False(t, pair.IsSome())
}

func TestByKey(t *testing.T) {
	m := map[int]string{1: "one", 2: "two"}
	pair := optpair.ByKey(1, m)
	v1, v2, ok := pair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 1, v1)
	assert.Equal(t, "one", v2)

	pair = optpair.ByKey(3, m)
	assert.False(t, pair.IsSome())
}

func TestMap(t *testing.T) {
	pair := optpair.Some(1, "one")
	mappedPair := optpair.Map(pair, func(v1 int, v2 string) (string, int) {
		return v2, v1
	})
	v1, v2, ok := mappedPair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, "one", v1)
	assert.Equal(t, 1, v2)

	nonePair := optpair.None[int, string]()
	mappedPair = optpair.Map(nonePair, func(v1 int, v2 string) (string, int) {
		return v2, v1
	})
	assert.False(t, mappedPair.IsSome())
}

func TestFlatMap(t *testing.T) {
	pair := optpair.Some(1, "one")
	flatMappedPair := optpair.FlatMap(pair, func(v1 int, v2 string) optpair.Pair[string, int] {
		return optpair.Some(v2, v1)
	})
	v1, v2, ok := flatMappedPair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, "one", v1)
	assert.Equal(t, 1, v2)

	nonePair := optpair.None[int, string]()
	flatMappedPair = optpair.FlatMap(nonePair, func(v1 int, v2 string) optpair.Pair[string, int] {
		return optpair.Some(v2, v1)
	})
	assert.False(t, flatMappedPair.IsSome())
}

func TestFilter(t *testing.T) {
	pair := optpair.Some(1, "one")
	filteredPair := optpair.Filter(pair, func(v1 int, _ string) bool {
		return v1 == 1
	})
	assert.True(t, filteredPair.IsSome())

	filteredPair = optpair.Filter(pair, func(v1 int, _ string) bool {
		return v1 == 2
	})
	assert.False(t, filteredPair.IsSome())
}

func TestJoinAnd(t *testing.T) {
	value1 := optval.Some(1)
	value2 := optval.Some("one")
	pair := optpair.JoinAnd(value1, value2)
	v1, v2, ok := pair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 1, v1)
	assert.Equal(t, "one", v2)

	value1 = optval.None[int]()
	pair = optpair.JoinAnd(value1, value2)
	assert.False(t, pair.IsSome())
}

func TestJoinOr(t *testing.T) {
	value1 := optval.Some(1)
	value2 := optval.Some("one")
	pair := optpair.JoinOr(value1, value2)
	v1, v2, ok := pair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 1, v1)
	assert.Equal(t, "one", v2)

	value1 = optval.None[int]()
	pair = optpair.JoinOr(value1, value2)
	v1, v2, ok = pair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 0, v1)
	assert.Equal(t, "one", v2)

	value1 = optval.None[int]()
	value2 = optval.None[string]()
	pair = optpair.JoinOr(value1, value2)
	assert.False(t, pair.IsSome())
}

func TestSplit(t *testing.T) {
	pair := optpair.Some(1, "one")
	value1, value2 := optpair.Split(pair)
	assert.True(t, value1.IsSome())
	assert.True(t, value2.IsSome())

	v1, ok := value1.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 1, v1)

	v2, ok := value2.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, "one", v2)

	pair = optpair.None[int, string]()
	value1, value2 = optpair.Split(pair)
	assert.False(t, value1.IsSome())
	assert.False(t, value2.IsSome())
}

func TestLeft(t *testing.T) {
	pair := optpair.Some(1, "one")
	left := optpair.Left(pair)
	assert.True(t, left.IsSome())
	v1, ok := left.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 1, v1)

	pair = optpair.None[int, string]()
	left = optpair.Left(pair)
	assert.False(t, left.IsSome())
}

func TestRight(t *testing.T) {
	pair := optpair.Some(1, "one")
	right := optpair.Right(pair)
	assert.True(t, right.IsSome())
	v2, ok := right.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, "one", v2)

	pair = optpair.None[int, string]()
	right = optpair.Right(pair)
	assert.False(t, right.IsSome())
}

func TestSwap(t *testing.T) {
	pair := optpair.Some(1, "one")
	swappedPair := optpair.Swap(pair)
	v1, v2, ok := swappedPair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, "one", v1)
	assert.Equal(t, 1, v2)

	pair = optpair.None[int, string]()
	swappedPair = optpair.Swap(pair)
	assert.False(t, swappedPair.IsSome())
}

func TestPair(t *testing.T) {
	pair := optpair.Some(1, "one")
	assert.True(t, pair.IsSome())
	assert.False(t, pair.IsNone())

	v1, v2, ok := pair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 1, v1)
	assert.Equal(t, "one", v2)

	otherPair := optpair.Some(2, "two")
	assert.Equal(t, pair, pair.Or(otherPair))

	nonePair := optpair.None[int, string]()
	assert.Equal(t, otherPair, nonePair.Or(otherPair))

	v1, v2 = pair.OrSome(2, "two")
	assert.Equal(t, 1, v1)
	assert.Equal(t, "one", v2)

	v1, v2 = nonePair.OrSome(2, "two")
	assert.Equal(t, 2, v1)
	assert.Equal(t, "two", v2)

	v1, v2 = pair.OrZero()
	assert.Equal(t, 1, v1)
	assert.Equal(t, "one", v2)

	v1, v2 = nonePair.OrZero()
	assert.Equal(t, 0, v1)
	assert.Equal(t, "", v2)

	assert.Equal(t, pair, pair.OrElse(func() optpair.Pair[int, string] {
		return otherPair
	}))

	assert.Equal(t, otherPair, nonePair.OrElse(func() optpair.Pair[int, string] {
		return otherPair
	}))

	v1, v2 = pair.OrElseSome(func() (int, string) {
		return 2, "two"
	})
	assert.Equal(t, 1, v1)
	assert.Equal(t, "one", v2)

	v1, v2 = nonePair.OrElseSome(func() (int, string) {
		return 2, "two"
	})
	assert.Equal(t, 2, v1)
	assert.Equal(t, "two", v2)

	pair.Reset()
	assert.False(t, pair.IsSome())

	pair = optpair.Some(1, "one")
	takenPair := pair.Take()
	assert.True(t, takenPair.IsSome())
	assert.False(t, pair.IsSome())

	pair = optpair.Some(1, "one")
	replacedPair := pair.Replace(2, "two")
	v1, v2, ok = replacedPair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 1, v1)
	assert.Equal(t, "one", v2)

	v1, v2, ok = pair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 2, v1)
	assert.Equal(t, "two", v2)
}

func TestUnwrapFilter(t *testing.T) {
	items := []optpair.Pair[int, string]{
		optpair.Some(1, "one"),
		optpair.None[int, string](),
		optpair.Some(2, "two"),
		optpair.None[int, string](),
		optpair.Some(3, "three"),
	}
	values := slices.Values(items)

	collected := map[int]string{}

	optpair.UnwrapFilter(values)(func(k int, v string) bool {
		collected[k] = v
		return true
	})

	assert.Equal(t, map[int]string{1: "one", 2: "two", 3: "three"}, collected)

	clear(collected)
	optpair.UnwrapFilter(values)(func(k int, v string) bool {
		collected[k] = v
		return k%2 != 0
	})

	assert.Equal(t, map[int]string{1: "one", 2: "two"}, collected)
}

func TestCompare(t *testing.T) {
	assert.Equal(t, 0, optpair.Compare(
		optpair.Some(42, "forty-two"),
		optpair.Some(42, "forty-two"),
	))
	assert.Equal(t, 1, optpair.Compare(
		optpair.Some(42, "two"),
		optpair.Some(42, "three"),
	))
	assert.Equal(t, -1, optpair.Compare(
		optpair.Some(42, "three"),
		optpair.Some(42, "two"),
	))
	assert.Equal(t, -1, optpair.Compare(
		optpair.Some(42, "some"),
		optpair.Some(43, "some"),
	))
	assert.Equal(t, 1, optpair.Compare(
		optpair.Some(42, "some"),
		optpair.Some(41, "some"),
	))
	assert.Equal(t, -1, optpair.Compare(
		optpair.Some(42, "some"),
		optpair.None[int, string](),
	))
	assert.Equal(t, 1, optpair.Compare(
		optpair.None[int, string](),
		optpair.Some(42, "some"),
	))
	assert.Equal(t, 0, optpair.Compare(
		optpair.None[int, string](),
		optpair.None[int, string](),
	))
}

func TestLess(t *testing.T) {
	assert.False(t, optpair.Less(
		optpair.Some(42, "forty-two"),
		optpair.Some(42, "forty-two"),
	))
	assert.True(t, optpair.Less(
		optpair.Some(42, "three"),
		optpair.Some(42, "two"),
	))
	assert.True(t, optpair.Less(
		optpair.Some(42, "some"),
		optpair.Some(43, "some"),
	))
	assert.False(t, optpair.Less(
		optpair.Some(42, "some"),
		optpair.Some(41, "some"),
	))
	assert.True(t, optpair.Less(
		optpair.Some(42, "some"),
		optpair.None[int, string](),
	))
	assert.False(t, optpair.Less(
		optpair.None[int, string](),
		optpair.Some(42, "some"),
	))
	assert.False(t, optpair.Less(
		optpair.None[int, string](),
		optpair.None[int, string](),
	))
}

func TestIsSome(t *testing.T) {
	assert.True(t, optpair.IsSome(optpair.Some(42, "forty-two")))
	assert.False(t, optpair.IsSome(optpair.None[int, string]()))
}

func TestIsNone(t *testing.T) {
	assert.False(t, optpair.IsNone(optpair.Some(42, "forty-two")))
	assert.True(t, optpair.IsNone(optpair.None[int, string]()))
}

func TestOr(t *testing.T) {
	pair := optpair.Some(1, "one")
	otherPair := optpair.Some(2, "two")
	assert.Equal(t, pair, optpair.Or(pair))
	assert.Equal(t, pair, optpair.Or(pair, otherPair))

	nonePair := optpair.None[int, string]()
	assert.Equal(t, otherPair, optpair.Or(nonePair, otherPair))
	assert.Equal(t, nonePair, optpair.Or(nonePair, nonePair))
}

func TestOrZero(t *testing.T) {
	group := func(v ...any) []any {
		return v
	}

	pair := optpair.Some(1, "one")
	assert.Equal(t, group(1, "one"), group(optpair.OrZero(pair)))

	nonePair := optpair.None[int, string]()
	assert.Equal(t, group(0, ""), group(optpair.OrZero(nonePair)))
}

func TestUnwrap(t *testing.T) {
	group := func(v ...any) []any {
		return v
	}

	pair := optpair.Some(1, "one")
	assert.Equal(t, group(1, "one", true), group(optpair.Unwrap(pair)))

	nonePair := optpair.None[int, string]()
	assert.Equal(t, group(0, "", false), group(optpair.Unwrap(nonePair)))
}
