package optional_test

import (
	"slices"
	"testing"

	"github.com/pamburus/go-mod/optional"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	v := optional.New(42, true)
	assert.True(t, v.IsSome())
	assert.Equal(t, 42, v.OrZero())

	v = optional.New(42, false)
	assert.True(t, v.IsNone())
}

func TestSome(t *testing.T) {
	v := optional.Some(42)
	assert.True(t, v.IsSome())
	assert.Equal(t, 42, v.OrZero())
}

func TestNone(t *testing.T) {
	v := optional.None[int]()
	assert.True(t, v.IsNone())
}

func TestValueByKey(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	v := optional.ValueByKey("a", m)
	assert.True(t, v.IsSome())
	assert.Equal(t, 1, v.OrZero())

	v = optional.ValueByKey("c", m)
	assert.True(t, v.IsNone())
}

func TestKey(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	v := optional.Key("a", m)
	assert.True(t, v.IsSome())
	assert.Equal(t, "a", v.OrZero())

	v = optional.Key("c", m)
	assert.True(t, v.IsNone())
}

func TestFromPtr(t *testing.T) {
	val := 42
	v := optional.FromPtr(&val)
	assert.True(t, v.IsSome())
	assert.Equal(t, 42, v.OrZero())

	v = optional.FromPtr((*int)(nil))
	assert.True(t, v.IsNone())
}

func TestMap(t *testing.T) {
	v := optional.Some(42)
	u := optional.Map(v, func(int) string {
		return "value"
	})
	assert.True(t, u.IsSome())
	assert.Equal(t, "value", u.OrZero())

	v = optional.None[int]()
	u = optional.Map(v, func(int) string {
		return "value"
	})
	assert.True(t, u.IsNone())
}

func TestFlatMap(t *testing.T) {
	v := optional.Some(42)
	u := optional.FlatMap(v, func(int) optional.Value[string] {
		return optional.Some("value")
	})
	assert.True(t, u.IsSome())
	assert.Equal(t, "value", u.OrZero())

	v = optional.None[int]()
	u = optional.FlatMap(v, func(int) optional.Value[string] {
		return optional.Some("value")
	})
	assert.True(t, u.IsNone())
}

func TestFilter(t *testing.T) {
	v := optional.Some(42)
	u := optional.Filter(v, func(i int) bool {
		return i > 40
	})
	assert.True(t, u.IsSome())
	assert.Equal(t, 42, u.OrZero())

	u = optional.Filter(v, func(i int) bool {
		return i < 40
	})
	assert.True(t, u.IsNone())
}

func TestCollect(t *testing.T) {
	values := []optional.Value[int]{optional.Some(1), optional.None[int](), optional.Some(2)}
	collected := optional.Collect(values...)
	assert.Equal(t, []int{1, 2}, collected)
}

func TestOr(t *testing.T) {
	v := optional.Some(42)
	u := optional.None[int]()
	assert.Equal(t, 42, v.Or(u).OrZero())
	assert.Equal(t, 42, u.Or(v).OrZero())
}

func TestOrSome(t *testing.T) {
	v := optional.Some(42)
	assert.Equal(t, 42, v.OrSome(0))

	u := optional.None[int]()
	assert.Equal(t, 0, u.OrSome(0))
}

func TestOrElse(t *testing.T) {
	v := optional.Some(42)
	assert.Equal(t, 42, v.OrElse(func() optional.Value[int] {
		return optional.Some(0)
	}).OrZero())

	u := optional.None[int]()
	assert.Equal(t, 0, u.OrElse(func() optional.Value[int] {
		return optional.Some(0)
	}).OrZero())
}

func TestOrElseSome(t *testing.T) {
	v := optional.Some(42)
	assert.Equal(t, 42, v.OrElseSome(func() int {
		return 0
	}))

	u := optional.None[int]()
	assert.Equal(t, 0, u.OrElseSome(func() int {
		return 0
	}))
}

func TestReset(t *testing.T) {
	v := optional.Some(42)
	v.Reset()
	assert.True(t, v.IsNone())
}

func TestTake(t *testing.T) {
	v := optional.Some(42)
	taken := v.Take()
	assert.True(t, taken.IsSome())
	assert.True(t, v.IsNone())
}

func TestReplace(t *testing.T) {
	v := optional.Some(42)
	replaced := v.Replace(100)
	assert.True(t, replaced.IsSome())
	assert.Equal(t, 42, replaced.OrZero())
	assert.Equal(t, 100, v.OrZero())
}

func TestCopyPtr(t *testing.T) {
	v := optional.Some(42)
	ptr := v.CopyPtr()
	assert.NotNil(t, ptr)
	assert.Equal(t, 42, *ptr)

	u := optional.None[int]()
	ptr = u.CopyPtr()
	assert.Nil(t, ptr)
}

func TestMapFromPtr(t *testing.T) {
	val := 42
	v := optional.MapFromPtr(&val, func(int) string {
		return "value"
	})
	assert.True(t, v.IsSome())
	assert.Equal(t, "value", v.OrZero())

	v = optional.MapFromPtr((*int)(nil), func(int) string {
		return "value"
	})
	assert.True(t, v.IsNone())
}

func TestFlatten(t *testing.T) {
	v := optional.Some(optional.Some(42))
	u := optional.Flatten(v)
	assert.True(t, u.IsSome())
	assert.Equal(t, 42, u.OrZero())

	v = optional.Some(optional.None[int]())
	u = optional.Flatten(v)
	assert.True(t, u.IsNone())

	v = optional.None[optional.Value[int]]()
	u = optional.Flatten(v)
	assert.True(t, u.IsNone())
}

func TestSomeOnly(t *testing.T) {
	items := []optional.Value[int]{
		optional.Some(1),
		optional.None[int](),
		optional.Some(2),
		optional.None[int](),
		optional.Some(3),
	}
	values := slices.Values(items)

	collected := []int{}
	optional.SomeOnly(values)(func(v int) bool {
		collected = append(collected, v)
		return true
	})

	assert.Equal(t, []int{1, 2, 3}, collected)

	collected = collected[:0]
	optional.SomeOnly(values)(func(v int) bool {
		collected = append(collected, v)
		return v%2 != 0
	})

	assert.Equal(t, []int{1, 2}, collected)
}
