package optional_test

import (
	"slices"
	"testing"

	"github.com/pamburus/go-mod/optional"
	"github.com/stretchr/testify/assert"
)

func TestNewPair(t *testing.T) {
	pair := optional.NewPair(1, "one", true)
	v1, v2, ok := pair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 1, v1)
	assert.Equal(t, "one", v2)

	pair = optional.NewPair(1, "one", false)
	assert.False(t, pair.IsSome())
}

func TestSomePair(t *testing.T) {
	pair := optional.SomePair(1, "one")
	v1, v2, ok := pair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 1, v1)
	assert.Equal(t, "one", v2)
}

func TestNonePair(t *testing.T) {
	pair := optional.NonePair[int, string]()
	assert.False(t, pair.IsSome())
}

func TestPairByKey(t *testing.T) {
	m := map[int]string{1: "one", 2: "two"}
	pair := optional.PairByKey(1, m)
	v1, v2, ok := pair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 1, v1)
	assert.Equal(t, "one", v2)

	pair = optional.PairByKey(3, m)
	assert.False(t, pair.IsSome())
}

func TestMapPair(t *testing.T) {
	pair := optional.SomePair(1, "one")
	mappedPair := optional.MapPair(pair, func(v1 int, v2 string) (string, int) {
		return v2, v1
	})
	v1, v2, ok := mappedPair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, "one", v1)
	assert.Equal(t, 1, v2)

	nonePair := optional.NonePair[int, string]()
	mappedPair = optional.MapPair(nonePair, func(v1 int, v2 string) (string, int) {
		return v2, v1
	})
	assert.False(t, mappedPair.IsSome())
}

func TestFlatMapPair(t *testing.T) {
	pair := optional.SomePair(1, "one")
	flatMappedPair := optional.FlatMapPair(pair, func(v1 int, v2 string) optional.Pair[string, int] {
		return optional.SomePair(v2, v1)
	})
	v1, v2, ok := flatMappedPair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, "one", v1)
	assert.Equal(t, 1, v2)

	nonePair := optional.NonePair[int, string]()
	flatMappedPair = optional.FlatMapPair(nonePair, func(v1 int, v2 string) optional.Pair[string, int] {
		return optional.SomePair(v2, v1)
	})
	assert.False(t, flatMappedPair.IsSome())
}

func TestFilterPair(t *testing.T) {
	pair := optional.SomePair(1, "one")
	filteredPair := optional.FilterPair(pair, func(v1 int, _ string) bool {
		return v1 == 1
	})
	assert.True(t, filteredPair.IsSome())

	filteredPair = optional.FilterPair(pair, func(v1 int, _ string) bool {
		return v1 == 2
	})
	assert.False(t, filteredPair.IsSome())
}

func TestBoth(t *testing.T) {
	value1 := optional.Some(1)
	value2 := optional.Some("one")
	pair := optional.Both(value1, value2)
	v1, v2, ok := pair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 1, v1)
	assert.Equal(t, "one", v2)

	value1 = optional.None[int]()
	pair = optional.Both(value1, value2)
	assert.False(t, pair.IsSome())
}

func TestLeft(t *testing.T) {
	pair := optional.SomePair(1, "one")
	left := optional.Left(pair)
	assert.True(t, left.IsSome())
	v1, ok := left.Unwrap()
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, v1)

	pair = optional.NonePair[int, string]()
	left = optional.Left(pair)
	assert.False(t, left.IsSome())
}

func TestRight(t *testing.T) {
	pair := optional.SomePair(1, "one")
	right := optional.Right(pair)
	assert.True(t, right.IsSome())
	v2, ok := right.Unwrap()
	assert.Equal(t, true, ok)
	assert.Equal(t, "one", v2)

	pair = optional.NonePair[int, string]()
	right = optional.Right(pair)
	assert.False(t, right.IsSome())
}

func TestPairMethods(t *testing.T) {
	pair := optional.SomePair(1, "one")
	assert.True(t, pair.IsSome())
	assert.False(t, pair.IsNone())

	v1, v2, ok := pair.Unwrap()
	assert.True(t, ok)
	assert.Equal(t, 1, v1)
	assert.Equal(t, "one", v2)

	otherPair := optional.SomePair(2, "two")
	assert.Equal(t, pair, pair.Or(otherPair))

	nonePair := optional.NonePair[int, string]()
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

	assert.Equal(t, pair, pair.OrElse(func() optional.Pair[int, string] {
		return otherPair
	}))

	assert.Equal(t, otherPair, nonePair.OrElse(func() optional.Pair[int, string] {
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

	pair = optional.SomePair(1, "one")
	takenPair := pair.Take()
	assert.True(t, takenPair.IsSome())
	assert.False(t, pair.IsSome())

	pair = optional.SomePair(1, "one")
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

func TestSomePairOnly(t *testing.T) {
	items := []optional.Pair[int, string]{
		optional.SomePair(1, "one"),
		optional.NonePair[int, string](),
		optional.SomePair(2, "two"),
		optional.NonePair[int, string](),
		optional.SomePair(3, "three"),
	}
	values := slices.Values(items)

	collected := map[int]string{}
	optional.SomePairOnly(values)(func(k int, v string) bool {
		collected[k] = v
		return true
	})

	assert.Equal(t, map[int]string{1: "one", 2: "two", 3: "three"}, collected)

	clear(collected)
	optional.SomePairOnly(values)(func(k int, v string) bool {
		collected[k] = v
		return k%2 != 0
	})

	assert.Equal(t, map[int]string{1: "one", 2: "two"}, collected)
}
