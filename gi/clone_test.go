package gi_test

import (
	"net/http"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
)

func TestCloned(t *testing.T) {
	headers := []http.Header{
		{"a": {"1"}},
		{"b": {"2"}},
		{"c": {"3"}},
	}

	result := slices.Collect(gi.Cloned(slices.Values(headers)))

	headers[0].Set("a", "4")

	expected := []http.Header{
		{"a": {"1"}},
		{"b": {"2"}},
		{"c": {"3"}},
	}
	assert.Equal(t, expected, result)
}
