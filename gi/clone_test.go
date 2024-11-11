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
		{"A": {"1"}},
		{"B": {"2"}},
		{"C": {"3"}},
	}

	result := slices.Collect(gi.Cloned(slices.Values(headers)))

	headers[0].Set("A", "4")

	expected := []http.Header{
		{"A": {"1"}},
		{"B": {"2"}},
		{"C": {"3"}},
	}
	assert.Equal(t, expected, result)
}
