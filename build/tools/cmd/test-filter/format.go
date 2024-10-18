package main

import "regexp"

var reformatSettings = []reformatSetting{
	{
		pattern: regexp.MustCompile(`^(?P<result>\?)\s+(?P<module>[\w\./\-]+)\s+(?P<notice>\[[\w\s]+\])\s*$`),
		format:  "%-7[1]s %-63[2]s         %[3]s",
	},
	{
		pattern: regexp.MustCompile(`^(?P<result>ok|fail|\?)\s+(?P<module>[\w\./\-]+)\s+(?P<duration>[\d\.]+s)\s+coverage:\s+(?P<coverage>[\d\.]+% of statements)\s*$`),
		format:  "%-7[1]s %-63[2]s %7[3]s coverage: %20[4]s",
	},
}

type reformatSetting struct {
	pattern *regexp.Regexp
	format  string
}
