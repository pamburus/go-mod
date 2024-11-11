package result_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pamburus/go-mod/result"
)

func TestWrapUnwrapPanic(t *testing.T) {
	err := result.WrapPanic(42)
	require.Error(t, err)
	require.NoError(t, errors.Unwrap(err))

	val, ok := result.UnwrapPanic(err)
	require.True(t, ok)
	assert.Equal(t, 42, val)
	assert.Equal(t, "panic: 42", err.Error())

	val, ok = result.UnwrapPanic(assert.AnError)
	require.False(t, ok)
	assert.Nil(t, val)

	err = result.WrapPanic(assert.AnError)
	require.Error(t, err)

	val, ok = result.UnwrapPanic(err)
	require.True(t, ok)
	assert.Equal(t, assert.AnError, val)
	assert.Equal(t, "panic: "+assert.AnError.Error(), err.Error())
	require.ErrorIs(t, err, assert.AnError)
}

func TestRecallPanic(t *testing.T) {
	err := result.WrapPanic(42)
	assert.PanicsWithValue(t, 42, func() { //nolint:wsl // err is used here
		_ = result.RecallPanic(err)
	})

	err = result.WrapPanic(assert.AnError)
	assert.PanicsWithError(t, assert.AnError.Error(), func() {
		_ = result.RecallPanic(err)
	})

	err = errors.New("error")
	assert.Equal(t, err, result.RecallPanic(err))
}
