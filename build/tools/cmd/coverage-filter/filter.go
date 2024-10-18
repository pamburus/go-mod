package main

import "regexp"

var verboseFilter = patternFilter{
	{
		positive: regexp.MustCompile(`\b100\.0%$`),
	},
}

// ---

type patternFilter []patternPair

func (f patternFilter) match(line string) bool {
	for _, pair := range f {
		if pair.match(line) {
			return true
		}
	}

	return false
}

// ---

type patternPair struct {
	negative *regexp.Regexp
	positive *regexp.Regexp
}

func (f patternPair) match(line string) bool {
	return f.matchPositive(line) && !f.matchNegative(line)
}

func (f patternPair) matchPositive(line string) bool {
	return f.positive == nil || f.positive.MatchString(line)
}

func (f patternPair) matchNegative(line string) bool {
	return f.negative != nil && f.negative.MatchString(line)
}
