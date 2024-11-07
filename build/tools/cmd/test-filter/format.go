package main

import "regexp"

var reformatSettings = []reformatSetting{
	{
		pattern: regexp.MustCompile(`^(?P<result>ok|\?)?\s+(?P<module>[\w\./\-]+)\s+(?:(?P<duration>\d+\.\d+s)\s+)?(?:(?:(?P<covprefix>coverage:)\s+(?P<coverage>\d+\.\d+%)(?P<covsuffix> of statements(?: in [\w\./\-]+)?))|(?P<notice>\[no test files\]))\s*$`),
		format:  "%-7[1]s %-79[2]s %7[3]s %[7]s%[4]s%7[5]s%[6]s",
	},
}

type reformatSetting struct {
	pattern *regexp.Regexp
	format  string
}
