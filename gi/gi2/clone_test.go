package gi2_test

import (
	"maps"
	"net/http"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestCloned(t *testing.T) {
	headers := []http.Header{
		{"A": {"1"}},
		{"B": {"2"}},
		{"C": {"3"}},
	}

	toPair := func(h http.Header) (http.Header, http.Header) {
		return h, h
	}

	cloned := gi2.Cloned(helpers.ToPairs(slices.Values(headers), toPair))
	result := slices.Collect(helpers.FlattenPairs(cloned))

	headers[0].Set("A", "6")
	headers[1].Set("B", "4")

	expected := []http.Header{
		{"A": {"1"}},
		{"A": {"1"}},
		{"B": {"2"}},
		{"B": {"2"}},
		{"C": {"3"}},
		{"C": {"3"}},
	}
	assert.Equal(t, expected, result)
}

func TestClonedRight(t *testing.T) {
	headers := []http.Header{
		{"A": {"1"}},
		{"B": {"2"}},
		{"C": {"3"}},
	}

	result := maps.Collect(gi2.ClonedRight(slices.All(headers)))

	headers[1].Set("B", "4")

	expected := map[int]http.Header{
		0: {"A": {"1"}},
		1: {"B": {"2"}},
		2: {"C": {"3"}},
	}
	assert.Equal(t, expected, result)
}

func TestClonedValues(t *testing.T) {
	headers := []http.Header{
		{"A": {"1"}},
		{"B": {"2"}},
		{"C": {"3"}},
	}

	result := maps.Collect(gi2.ClonedValues(slices.All(headers)))

	headers[1].Set("B", "4")

	expected := map[int]http.Header{
		0: {"A": {"1"}},
		1: {"B": {"2"}},
		2: {"C": {"3"}},
	}
	assert.Equal(t, expected, result)
}

func TestClonedLeft(t *testing.T) {
	headers := []http.Header{
		{"A": {"1"}},
		{"B": {"2"}},
		{"C": {"3"}},
	}

	cloned := gi2.ClonedLeft(helpers.Swap(slices.All(headers)))
	result := maps.Collect(helpers.Swap(cloned))

	headers[1].Set("B", "4")

	expected := map[int]http.Header{
		0: {"A": {"1"}},
		1: {"B": {"2"}},
		2: {"C": {"3"}},
	}
	assert.Equal(t, expected, result)
}

func TestClonedKeys(t *testing.T) {
	headers := []http.Header{
		{"A": {"1"}},
		{"B": {"2"}},
		{"C": {"3"}},
	}

	cloned := gi2.ClonedKeys(helpers.Swap(slices.All(headers)))
	result := maps.Collect(helpers.Swap(cloned))

	headers[1].Set("B", "4")

	expected := map[int]http.Header{
		0: {"A": {"1"}},
		1: {"B": {"2"}},
		2: {"C": {"3"}},
	}
	assert.Equal(t, expected, result)
}
