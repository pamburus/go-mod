package colorization

import (
	"os"

	"github.com/mattn/go-isatty"
)

func Enabled(flag string, stream *os.File) bool {
	switch flag {
	default:
		fallthrough
	case "auto":
		return isatty.IsTerminal(stream.Fd())
	case "always":
		return true
	case "never":
		return false
	}
}
