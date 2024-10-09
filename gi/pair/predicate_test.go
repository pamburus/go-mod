package pair_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/pair"
	"github.com/pamburus/go-mod/optional/optpair"
)

func TestEqual(t *testing.T) {
	t.Run("SomeEqual", func(t *testing.T) {
		pairs := slices.All([]int{10, 20, 30, 40, 50, 60})
		result := gi.FindPair(pairs, pair.Equal(2, 30))
		assert.Equal(t, optpair.Some(2, 30), result)
	})

	t.Run("NoneEqual", func(t *testing.T) {
		pairs := slices.All([]int{10, 20, 30, 40, 50, 60})
		result := gi.FindPair(pairs, pair.Equal(2, 70))
		assert.Equal(t, optpair.None[int, int](), result)
	})
}

func TestNotEqual(t *testing.T) {
	t.Run("SomeNotEqual", func(t *testing.T) {
		pairs := slices.All([]int{10, 20, 30, 40, 50, 60})
		result := gi.FindPair(pairs, pair.NotEqual(2, 70))
		assert.Equal(t, optpair.Some(0, 10), result)
	})

	t.Run("NoneNotEqual", func(t *testing.T) {
		pairs := slices.All([]int{10})
		result := gi.FindPair(pairs, pair.NotEqual(0, 10))
		assert.Equal(t, optpair.None[int, int](), result)
	})
}

func TestLess(t *testing.T) {
	t.Run("SomeLess", func(t *testing.T) {
		pairs := slices.All([]int{70, 50, 30, 42, 50, 60})
		result := gi.FindPair(pairs, pair.Less(3, 50))
		assert.Equal(t, optpair.Some(0, 70), result)
	})

	t.Run("NoneLess", func(t *testing.T) {
		pairs := slices.All([]int{10, 20, 30, 40, 50, 60})
		result := gi.FindPair(pairs, pair.Less(0, 10))
		assert.Equal(t, optpair.None[int, int](), result)
	})
}

func TestLessOrEqual(t *testing.T) {
	t.Run("SomeLessOrEqual", func(t *testing.T) {
		pairs := slices.All([]int{70, 50, 30, 42, 50, 60})
		result := gi.FindPair(pairs, pair.LessOrEqual(3, 50))
		assert.Equal(t, optpair.Some(0, 70), result)
	})

	t.Run("SomeEqual", func(t *testing.T) {
		pairs := slices.All([]int{70, 50, 30, 42, 50, 60})
		result := gi.FindPair(pairs, pair.LessOrEqual(0, 70))
		assert.Equal(t, optpair.Some(0, 70), result)
	})

	t.Run("NoneLessOrEqual", func(t *testing.T) {
		pairs := slices.All([]int{10, 20, 30, 40, 50, 60})
		result := gi.FindPair(pairs, pair.LessOrEqual(0, 9))
		assert.Equal(t, optpair.None[int, int](), result)
	})
}

func TestGreater(t *testing.T) {
	t.Run("SomeGreater", func(t *testing.T) {
		pairs := slices.All([]int{10, 20, 30, 40, 50, 60})
		result := gi.FindPair(pairs, pair.Greater(3, 42))
		assert.Equal(t, optpair.Some(4, 50), result)
	})

	t.Run("NoneGreater", func(t *testing.T) {
		pairs := slices.All([]int{10, 20, 30, 40, 50, 60})
		result := gi.FindPair(pairs, pair.Greater(5, 60))
		assert.Equal(t, optpair.None[int, int](), result)
	})
}

func TestGreaterOrEqual(t *testing.T) {
	t.Run("SomeGreaterOrEqual", func(t *testing.T) {
		pairs := slices.All([]int{10, 20, 30, 40, 50, 60})
		result := gi.FindPair(pairs, pair.GreaterOrEqual(3, 42))
		assert.Equal(t, optpair.Some(4, 50), result)
	})

	t.Run("SomeEqual", func(t *testing.T) {
		pairs := slices.All([]int{10, 20, 30, 40, 50, 60})
		result := gi.FindPair(pairs, pair.GreaterOrEqual(4, 50))
		assert.Equal(t, optpair.Some(4, 50), result)
	})

	t.Run("NoneGreaterOrEqual", func(t *testing.T) {
		pairs := slices.All([]int{10, 20, 30, 40, 50, 60})
		result := gi.FindPair(pairs, pair.GreaterOrEqual(5, 61))
		assert.Equal(t, optpair.None[int, int](), result)
	})
}

func TestAnd(t *testing.T) {
	t.Run("SomeAnd", func(t *testing.T) {
		pairs := slices.All([]int{10, 20, 30, 40, 50, 60})
		result := gi.FindPair(pairs, pair.And(pair.Greater(3, 42), pair.Less(4, 51)))
		assert.Equal(t, optpair.Some(4, 50), result)
	})

	t.Run("NoneAnd", func(t *testing.T) {
		pairs := slices.All([]int{10, 20, 30, 40, 50, 60})
		result := gi.FindPair(pairs, pair.And(pair.Greater(3, 42), pair.Less(4, 40)))
		assert.Equal(t, optpair.None[int, int](), result)
	})
}

func TestOr(t *testing.T) {
	t.Run("SomeOr", func(t *testing.T) {
		pairs := slices.All([]int{10, 20, 30, 40, 50, 60})
		result := gi.FindPair(pairs, pair.Or(pair.Greater(3, 42), pair.Less(-1, 51)))
		assert.Equal(t, optpair.Some(4, 50), result)
	})

	t.Run("SomeFirstOr", func(t *testing.T) {
		pairs := slices.All([]int{10, 20, 30, 40, 50, 60})
		result := gi.FindPair(pairs, pair.Or(pair.Greater(3, 42), pair.Less(4, 40)))
		assert.Equal(t, optpair.Some(0, 10), result)
	})

	t.Run("NoneOr", func(t *testing.T) {
		pairs := slices.All([]int{10, 20, 30, 40, 50, 60})
		result := gi.FindPair(pairs, pair.Or(pair.Greater(5, 60), pair.Less(0, 10)))
		assert.Equal(t, optpair.None[int, int](), result)
	})
}
