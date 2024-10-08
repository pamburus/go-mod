package optval_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/optional/optval"
)

func TestNew(t *testing.T) {
	v := optval.New(42, true)
	assert.True(t, v.IsSome())
	assert.Equal(t, 42, v.OrZero())

	v = optval.New(42, false)
	assert.True(t, v.IsNone())
}

func TestSome(t *testing.T) {
	v := optval.Some(42)
	assert.True(t, v.IsSome())
	assert.Equal(t, 42, v.OrZero())
}

func TestNone(t *testing.T) {
	v := optval.None[int]()
	assert.True(t, v.IsNone())
}

func TestValueByKey(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	v := optval.ByKey("a", m)
	assert.True(t, v.IsSome())
	assert.Equal(t, 1, v.OrZero())

	v = optval.ByKey("c", m)
	assert.True(t, v.IsNone())
}

func TestKey(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	v := optval.Key("a", m)
	assert.True(t, v.IsSome())
	assert.Equal(t, "a", v.OrZero())

	v = optval.Key("c", m)
	assert.True(t, v.IsNone())
}

func TestFromPtr(t *testing.T) {
	val := 42
	v := optval.FromPtr(&val)
	assert.True(t, v.IsSome())
	assert.Equal(t, 42, v.OrZero())

	v = optval.FromPtr((*int)(nil))
	assert.True(t, v.IsNone())
}

func TestMap(t *testing.T) {
	v := optval.Some(42)
	u := optval.Map(v, func(int) string {
		return "value"
	})
	assert.True(t, u.IsSome())
	assert.Equal(t, "value", u.OrZero())

	v = optval.None[int]()
	u = optval.Map(v, func(int) string {
		return "value"
	})
	assert.True(t, u.IsNone())
}

func TestFlatMap(t *testing.T) {
	v := optval.Some(42)
	u := optval.FlatMap(v, func(int) optval.Value[string] {
		return optval.Some("value")
	})
	assert.True(t, u.IsSome())
	assert.Equal(t, "value", u.OrZero())

	v = optval.None[int]()
	u = optval.FlatMap(v, func(int) optval.Value[string] {
		return optval.Some("value")
	})
	assert.True(t, u.IsNone())
}

func TestFilter(t *testing.T) {
	v := optval.Some(42)
	u := optval.Filter(v, func(i int) bool {
		return i > 40
	})
	assert.True(t, u.IsSome())
	assert.Equal(t, 42, u.OrZero())

	u = optval.Filter(v, func(i int) bool {
		return i < 40
	})
	assert.True(t, u.IsNone())
}

func TestCollect(t *testing.T) {
	values := []optval.Value[int]{optval.Some(1), optval.None[int](), optval.Some(2)}
	collected := optval.Collect(values...)
	assert.Equal(t, []int{1, 2}, collected)
}

func TestOr(t *testing.T) {
	v := optval.Some(42)
	u := optval.None[int]()
	assert.Equal(t, 42, v.Or(u).OrZero())
	assert.Equal(t, 42, u.Or(v).OrZero())
}

func TestOrSome(t *testing.T) {
	v := optval.Some(42)
	assert.Equal(t, 42, v.OrSome(0))

	u := optval.None[int]()
	assert.Equal(t, 0, u.OrSome(0))
}

func TestOrElse(t *testing.T) {
	v := optval.Some(42)
	assert.Equal(t, 42, v.OrElse(func() optval.Value[int] {
		return optval.Some(0)
	}).OrZero())

	u := optval.None[int]()
	assert.Equal(t, 0, u.OrElse(func() optval.Value[int] {
		return optval.Some(0)
	}).OrZero())
}

func TestOrElseSome(t *testing.T) {
	v := optval.Some(42)
	assert.Equal(t, 42, v.OrElseSome(func() int {
		return 0
	}))

	u := optval.None[int]()
	assert.Equal(t, 0, u.OrElseSome(func() int {
		return 0
	}))
}

func TestReset(t *testing.T) {
	v := optval.Some(42)
	v.Reset()
	assert.True(t, v.IsNone())
}

func TestTake(t *testing.T) {
	v := optval.Some(42)
	taken := v.Take()
	assert.True(t, taken.IsSome())
	assert.True(t, v.IsNone())
}

func TestReplace(t *testing.T) {
	v := optval.Some(42)
	replaced := v.Replace(100)
	assert.True(t, replaced.IsSome())
	assert.Equal(t, 42, replaced.OrZero())
	assert.Equal(t, 100, v.OrZero())
}

func TestCopyPtr(t *testing.T) {
	v := optval.Some(42)
	ptr := v.CopyPtr()
	assert.NotNil(t, ptr)
	assert.Equal(t, 42, *ptr)

	u := optval.None[int]()
	ptr = u.CopyPtr()
	assert.Nil(t, ptr)
}

func TestMapFromPtr(t *testing.T) {
	val := 42
	v := optval.MapFromPtr(&val, func(int) string {
		return "value"
	})
	assert.True(t, v.IsSome())
	assert.Equal(t, "value", v.OrZero())

	v = optval.MapFromPtr((*int)(nil), func(int) string {
		return "value"
	})
	assert.True(t, v.IsNone())
}

func TestFlatten(t *testing.T) {
	v := optval.Some(optval.Some(42))
	u := optval.Flatten(v)
	assert.True(t, u.IsSome())
	assert.Equal(t, 42, u.OrZero())

	v = optval.Some(optval.None[int]())
	u = optval.Flatten(v)
	assert.True(t, u.IsNone())

	v = optval.None[optval.Value[int]]()
	u = optval.Flatten(v)
	assert.True(t, u.IsNone())
}

func TestSomeOnly(t *testing.T) {
	items := []optval.Value[int]{
		optval.Some(1),
		optval.None[int](),
		optval.Some(2),
		optval.None[int](),
		optval.Some(3),
	}
	values := slices.Values(items)

	collected := []int{}
	optval.SomeOnly(values)(func(v int) bool {
		collected = append(collected, v)
		return true
	})

	assert.Equal(t, []int{1, 2, 3}, collected)

	collected = collected[:0]
	optval.SomeOnly(values)(func(v int) bool {
		collected = append(collected, v)
		return v%2 != 0
	})

	assert.Equal(t, []int{1, 2}, collected)
}