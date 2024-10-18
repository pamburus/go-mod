package cmpx_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/optional/internal/cmpx"
)

func TestIfElse(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		assert.Equal(t, 42, cmpx.IfElse(true, 42, 0))
	})

	t.Run("false", func(t *testing.T) {
		assert.Equal(t, 0, cmpx.IfElse(false, 42, 0))
	})
}
