package result_test

import (
	"errors"
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pamburus/go-mod/result"
)

func TestOk(t *testing.T) {
	assert.Equal(t,
		result.New(42, nil),
		result.Ok(42),
	)

	res := result.New(42, nil)
	assert.True(t, res.IsOk())
	assert.False(t, res.IsErr())
	assert.False(t, res.IsPanic())

	val, ok := res.Value()
	assert.True(t, ok)
	assert.Equal(t, 42, val)

	val, err := res.Unwrap()
	assert.Equal(t, 42, val)
	require.NoError(t, err)

	val, err = res.UnwrapNoPanic()
	assert.Equal(t, 42, val)
	require.NoError(t, err)

	assert.NoError(t, res.Err())

	val, err = result.Unwrap(res)
	assert.Equal(t, 42, val)
	require.NoError(t, err)

	val, err = result.UnwrapNoPanic(res)
	assert.Equal(t, 42, val)
	require.NoError(t, err)

	val, ok = result.Value(res)
	assert.True(t, ok)
	assert.Equal(t, 42, val)

	assert.True(t, result.IsOk(res))
	assert.False(t, result.IsErr(res))
	assert.False(t, result.IsPanic(res))
	require.NoError(t, result.Err(res))
}

func TestErr(t *testing.T) {
	assert.Equal(t,
		result.New(0, assert.AnError),
		result.NewErr[int](assert.AnError),
	)

	res := result.New(42, assert.AnError)
	assert.False(t, res.IsOk())
	assert.True(t, res.IsErr())
	assert.False(t, res.IsPanic())

	val, ok := res.Value()
	assert.False(t, ok)
	assert.Equal(t, 42, val)

	val, err := res.Unwrap()
	assert.Equal(t, 42, val)
	assert.Equal(t, err, assert.AnError)

	val, err = res.UnwrapNoPanic()
	assert.Equal(t, 42, val)
	assert.Equal(t, err, assert.AnError)

	assert.Equal(t, err, res.Err())
	assert.Equal(t, err, result.Err(res))
}

func TestGet(t *testing.T) {
	assert.Equal(t,
		result.New(42, nil),
		result.Get(func() (int, error) {
			return 42, nil
		}),
	)

	assert.Equal(t,
		result.New(42, assert.AnError),
		result.Get(func() (int, error) {
			return 42, assert.AnError
		}),
	)

	assert.Equal(t,
		result.NewErr[int](assert.AnError),
		result.Get(func() (int, error) {
			return 0, assert.AnError
		}),
	)

	assert.Equal(t,
		result.NewErr[int](result.WrapPanic(42)),
		result.Get(func() (int, error) {
			panic(42)
		}),
	)
}

func TestMap(t *testing.T) {
	t.Run("Ok", func(t *testing.T) {
		res := result.Ok(42)
		res = result.Map(res, func(v int) int {
			return v + 1
		})
		assert.Equal(t, result.Ok(43), res)
	})

	t.Run("Err", func(t *testing.T) {
		res := result.NewErr[int](assert.AnError)
		res = result.Map(res, func(v int) int {
			return v + 1
		})
		assert.Equal(t, result.NewErr[int](assert.AnError), res)
	})
}

func TestMapErr(t *testing.T) {
	t.Run("Ok", func(t *testing.T) {
		res := result.Ok(42)
		res = result.MapErr(res, func(error) error {
			return assert.AnError
		})
		assert.Equal(t, result.Ok(42), res)
	})

	t.Run("Err", func(t *testing.T) {
		res := result.New(42, assert.AnError)
		res = result.MapErr(res, func(e error) error {
			return fmt.Errorf("wrapped: %w", e)
		})
		assert.Equal(t, result.New(42, fmt.Errorf("wrapped: %w", assert.AnError)), res)
	})
}

func TestFlatten(t *testing.T) {
	t.Run("OkOk", func(t *testing.T) {
		res := result.Ok(result.Ok(42))
		assert.Equal(t, result.Ok(42), result.Flatten(res))
	})

	t.Run("OkErr", func(t *testing.T) {
		res := result.Ok(result.New(42, assert.AnError))
		assert.Equal(t, result.New(42, assert.AnError), result.Flatten(res))
	})

	t.Run("OkPanic", func(t *testing.T) {
		res := result.Ok(result.NewPanic[int](42))
		assert.Equal(t, result.NewPanic[int](42), result.Flatten(res))
	})

	t.Run("ErrOk", func(t *testing.T) {
		res := result.New(result.Ok(42), assert.AnError)
		assert.Equal(t, result.New(42, assert.AnError), result.Flatten(res))
	})

	t.Run("ErrErr", func(t *testing.T) {
		res := result.New(result.New(42, errors.New("another error")), assert.AnError)
		assert.Equal(t, result.New(42, assert.AnError), result.Flatten(res))
	})

	t.Run("ErrPanic", func(t *testing.T) {
		res := result.New(result.NewPanic[int](42), assert.AnError)
		assert.Equal(t, result.NewErr[int](assert.AnError), result.Flatten(res))
	})
}

func TestFlatMap(t *testing.T) {
	t.Run("OkOk", func(t *testing.T) {
		res := result.Ok(42)
		res = result.FlatMap(res, func(v int) (int, error) {
			return v + 1, nil
		})
		assert.Equal(t, result.Ok(43), res)
	})

	t.Run("OkErr", func(t *testing.T) {
		res := result.Ok(42)
		res = result.FlatMap(res, func(int) (int, error) {
			return 0, assert.AnError
		})
		assert.Equal(t, result.NewErr[int](assert.AnError), res)
	})

	t.Run("ErrOk", func(t *testing.T) {
		res := result.New(42, assert.AnError)
		res = result.FlatMap(res, func(v int) (int, error) {
			return v + 1, nil
		})
		assert.Equal(t, result.New(0, assert.AnError), res)
	})

	t.Run("ErrErr", func(t *testing.T) {
		res := result.New(42, assert.AnError)
		res = result.FlatMap(res, func(int) (int, error) {
			return 0, errors.New("another error")
		})
		assert.Equal(t, result.New(0, assert.AnError), res)
	})
}

func TestJoin(t *testing.T) {
	res := result.Join(
		result.Ok(42),
		result.New(43, assert.AnError),
		result.NewErr[int](errors.New("another error")),
		result.Ok(44),
		result.NewPanic[int](45),
	)

	assert.Equal(t,
		result.New(
			[]int{42, 44},
			errors.Join(
				assert.AnError,
				errors.New("another error"),
				result.WrapPanic(45),
			),
		),
		res,
	)
}

func TestJoinSeq(t *testing.T) {
	results := slices.Values([]result.Result[int]{
		result.Ok(42),
		result.New(43, assert.AnError),
		result.NewErr[int](errors.New("another error")),
		result.Ok(44),
		result.NewPanic[int](45),
	})

	res := result.JoinSeq(results)

	assert.Equal(t,
		result.New(
			[]int{42, 44},
			errors.Join(
				assert.AnError,
				errors.New("another error"),
				result.WrapPanic(45),
			),
		),
		res,
	)
}

func TestFromSeq2(t *testing.T) {
	seq := func(yield func(int, error) bool) {
		_ = yield(42, nil) &&
			yield(43, assert.AnError) &&
			yield(44, nil)
	}

	results := slices.Collect(result.FromSeq2(seq))

	assert.Equal(t,
		[]result.Result[int]{
			result.Ok(42),
			result.New(43, assert.AnError),
			result.Ok(44),
		},
		results,
	)
}

func TestUnwrapCollect(t *testing.T) {
	t.Run("WithErrors", func(t *testing.T) {
		seq := func(yield func(result.Result[int]) bool) {
			_ = yield(result.Ok(42)) &&
				yield(result.New(43, assert.AnError)) &&
				yield(result.Ok(44))
		}

		items, err := result.UnwrapCollect([]int(nil), seq)
		require.ErrorIs(t, err, assert.AnError)
		assert.Nil(t, items)
	})

	t.Run("WithoutErrors", func(t *testing.T) {
		seq := func(yield func(result.Result[int]) bool) {
			_ = yield(result.Ok(42)) &&
				yield(result.Ok(43)) &&
				yield(result.Ok(44))
		}

		items, err := result.UnwrapCollect([]int(nil), seq)
		require.NoError(t, err)
		assert.Equal(t, []int{42, 43, 44}, items)
	})
}
