package giop2_test

import (
	"testing"

	"github.com/pamburus/go-mod/gi/giop2"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, []any{3, 5}, group(giop2.Add(1, 2, 2, 3)))
		assert.Equal(t, []any{5, 6}, group(giop2.Add(3, 2, 2, 4)))
	})

	t.Run("Float", func(t *testing.T) {
		assert.Equal(t, []any{3.5, 5.5}, group(giop2.Add(1.5, 2.0, 2.0, 3.5)))
		assert.Equal(t, []any{7.0, 5.5}, group(giop2.Add(4.5, 2.0, 2.5, 3.5)))
	})
}

func TestSubtract(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, []any{-1, -1}, group(giop2.Subtract(1, 2, 2, 3)))
		assert.Equal(t, []any{1, -2}, group(giop2.Subtract(3, 2, 2, 4)))
	})

	t.Run("Float", func(t *testing.T) {
		assert.Equal(t, []any{-0.5, -1.5}, group(giop2.Subtract(1.5, 2.0, 2.0, 3.5)))
		assert.Equal(t, []any{2.0, -1.5}, group(giop2.Subtract(4.5, 2.0, 2.5, 3.5)))
	})
}
