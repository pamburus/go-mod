// Package envflag provides a way to read boolean flags from environment variables.
package envflag

import (
	"os"
	"strings"
)

// Get returns the value of the environment variable name as a boolean flag.
func Get(name string) bool {
	if val, ok := os.LookupEnv(name); ok {
		return enabled[strings.ToLower(val)]
	}

	return false
}

// ---

var enabled = map[string]bool{
	"on":   true,
	"yes":  true,
	"1":    true,
	"true": true,
}
