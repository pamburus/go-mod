package gi_test

import (
	"maps"
	"net/http"
	"slices"
	"testing"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
	"github.com/stretchr/testify/assert"
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

func TestClonedPairs(t *testing.T) {
	headers := []http.Header{
		{"a": {"1"}},
		{"b": {"2"}},
		{"c": {"3"}},
	}

	toPair := func(h http.Header) (http.Header, http.Header) {
		return h, h
	}

	cloned := gi.ClonedPairs(helpers.ToPairs(slices.Values(headers), toPair))
	result := slices.Collect(helpers.FlattenPairs(cloned))

	headers[0].Set("a", "6")
	headers[1].Set("b", "4")

	expected := []http.Header{
		{"a": {"1"}},
		{"a": {"1"}},
		{"b": {"2"}},
		{"b": {"2"}},
		{"c": {"3"}},
		{"c": {"3"}},
	}
	assert.Equal(t, expected, result)
}

func TestClonedRight(t *testing.T) {
	headers := []http.Header{
		{"a": {"1"}},
		{"b": {"2"}},
		{"c": {"3"}},
	}

	result := maps.Collect(gi.ClonedRight(slices.All(headers)))

	headers[1].Set("b", "4")

	expected := map[int]http.Header{
		0: {"a": {"1"}},
		1: {"b": {"2"}},
		2: {"c": {"3"}},
	}
	assert.Equal(t, expected, result)
}

func TestClonedValues(t *testing.T) {
	headers := []http.Header{
		{"a": {"1"}},
		{"b": {"2"}},
		{"c": {"3"}},
	}

	result := maps.Collect(gi.ClonedValues(slices.All(headers)))

	headers[1].Set("b", "4")

	expected := map[int]http.Header{
		0: {"a": {"1"}},
		1: {"b": {"2"}},
		2: {"c": {"3"}},
	}
	assert.Equal(t, expected, result)
}

func TestClonedLeft(t *testing.T) {
	headers := []http.Header{
		{"a": {"1"}},
		{"b": {"2"}},
		{"c": {"3"}},
	}

	cloned := gi.ClonedLeft(helpers.Swap(slices.All(headers)))
	result := maps.Collect(helpers.Swap(cloned))
	headers[1].Set("b", "4")

	expected := map[int]http.Header{
		0: {"a": {"1"}},
		1: {"b": {"2"}},
		2: {"c": {"3"}},
	}
	assert.Equal(t, expected, result)
}

func TestClonedKeys(t *testing.T) {
	headers := []http.Header{
		{"a": {"1"}},
		{"b": {"2"}},
		{"c": {"3"}},
	}

	cloned := gi.ClonedKeys(helpers.Swap(slices.All(headers)))
	result := maps.Collect(helpers.Swap(cloned))
	headers[1].Set("b", "4")

	expected := map[int]http.Header{
		0: {"a": {"1"}},
		1: {"b": {"2"}},
		2: {"c": {"3"}},
	}
	assert.Equal(t, expected, result)
}
