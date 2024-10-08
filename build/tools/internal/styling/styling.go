package styling

import (
	"regexp"

	"github.com/pamburus/go-ansi-esc/sgr"
)

// ---

func NewSetting(pattern string, style Style) Setting {
	return Setting{
		pattern: regexp.MustCompile(pattern),
		style:   style,
	}
}

func NewStyle(intro, outro string) Style {
	return Style{intro, outro}
}

func NewSequence(seq ...sgr.Command) string {
	return string(sgr.Sequence(seq).Bytes())
}

// ---

type Setting struct {
	pattern *regexp.Regexp
	style   Style
}

func (s Setting) Apply(line string) string {
	return s.pattern.ReplaceAllStringFunc(line, func(match string) string {
		return s.style.Apply(match)
	})
}

// ---

type Style struct {
	intro string
	outro string
}

func (s Style) Apply(line string) string {
	return s.intro + line + s.outro
}
