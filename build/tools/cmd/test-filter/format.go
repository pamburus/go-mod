package main

import "regexp"

var reformatSettings = []reformatSetting{
	{
		pattern: regexp.MustCompile(`^(?P<result>ok|\?)?\s+(?P<module>[\w\./\-]+)\s+(?:(?P<duration>\d+\.\d+s)\s+)?(?:(?:(?P<covprefix>coverage: )\s*(?P<coverage>(?:\d+\.\d+%)|(?:\[no statements\]))(?P<covsuffix>(?: of statements(?: in [\w\./\-]+)?)|(?:\s*\[no tests to run\]))?)|(?:\s*(?P<notice>(?:\s*\[no [a-z ]+\])+)))\s*$`),
		format:  "%-4[1]s\t%-79[2]s %7[3]s %[7]s%[4]s%6[5]s%[6]s",
	},
}

type reformatSetting struct {
	pattern *regexp.Regexp
	format  string
}
