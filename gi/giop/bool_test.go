package giop_test

import (
	"testing"

	"github.com/pamburus/go-mod/gi/giop"
	"github.com/stretchr/testify/assert"
)

func TestAnd(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		assert.False(t, giop.And(false, false))
		assert.False(t, giop.And(false, true))
		assert.False(t, giop.And(true, false))
		assert.True(t, giop.And(true, true))
	})
}

func TestOr(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		assert.False(t, giop.Or(false, false))
		assert.True(t, giop.Or(false, true))
		assert.True(t, giop.Or(true, false))
		assert.True(t, giop.Or(true, true))
	})
}

func TestXor(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		assert.False(t, giop.Xor(false, false))
		assert.True(t, giop.Xor(false, true))
		assert.True(t, giop.Xor(true, false))
		assert.False(t, giop.Xor(true, true))
	})
}
