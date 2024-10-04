package cmpx_test

import (
	"testing"

	"github.com/pamburus/go-mod/optional/internal/cmpx"
	"github.com/stretchr/testify/assert"
)

func TestIfElse(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		assert.Equal(t, 42, cmpx.IfElse(true, 42, 0))
	})

	t.Run("false", func(t *testing.T) {
		assert.Equal(t, 0, cmpx.IfElse(false, 42, 0))
	})
}
