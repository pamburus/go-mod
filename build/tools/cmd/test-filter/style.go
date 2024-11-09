package main

import (
	"github.com/pamburus/go-ansi-esc/sgr"
	"github.com/pamburus/go-mod/build/tools/internal/styling"
)

var styleSettings = []styling.Setting{
	styling.NewSetting(
		`^\s*--- PASS: [\w/]+\b`,
		styling.NewStyle(
			styling.NewSequence(),
			styling.NewSequence(),
		),
	),
	styling.NewSetting(
		`^\s*--- FAIL: .* \(\d+\.\d+s\)$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.BrightRed)),
			styling.NewSequence(sgr.ResetForegroundColor),
		),
	),
	styling.NewSetting(
		`^PASS$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Black), sgr.SetBackgroundColor(sgr.Green), sgr.SetBold)+"[",
			"]"+styling.NewSequence(sgr.ResetAll),
		),
	),
	styling.NewSetting(
		`^FAIL$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Black), sgr.SetBackgroundColor(sgr.Red), sgr.SetBold)+"[",
			"]"+styling.NewSequence(sgr.ResetAll),
		),
	),
	styling.NewSetting(
		`^FAIL\s+.*`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Red)),
			styling.NewSequence(sgr.ResetForegroundColor),
		),
	),
	styling.NewSetting(
		`^(?P<result>ok|\?)?\s+(?P<module>[\w\./\-]+)\s+(?P<duration>\d+\.\d+s)?\s+(?P<coverage>(?:coverage:\s+(?:(?:\d+\.\d+% of statements)|(?:\[no statements\]))))?(?:\s*\[no [a-z ]+\])*(?P<in>\s+in [\w\./\-]+)?\s*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetFaint),
			styling.NewSequence(sgr.ResetBoldAndFaint),
		),
	),
	styling.NewSetting(
		`\[no [a-z ]+\]`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetItalic, sgr.SetForegroundColor(sgr.BrightBlack)),
			styling.NewSequence(sgr.ResetItalic, sgr.ResetForegroundColor),
		),
	),
	styling.NewSetting(
		`^=== (?:[A-Z]+)\s.*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetFaint),
			styling.NewSequence(sgr.ResetBoldAndFaint),
		),
	),
	styling.NewSetting(
		`[^\s]+\.go(:\d+)?`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Cyan)),
			styling.NewSequence(sgr.ResetForegroundColor),
		),
	),
	styling.NewSetting(
		`\s*Error(?: Trace)?:.*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Red)),
			styling.NewSequence(sgr.ResetForegroundColor),
		),
	),
	styling.NewSetting(
		`\bexpected\s*: .*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Blue)),
			styling.NewSequence(sgr.ResetForegroundColor),
		),
	),
	styling.NewSetting(
		`\bactual\s*: .*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Red), sgr.SetUnderlined),
			styling.NewSequence(sgr.ResetForegroundColor, sgr.ResetAllUnderlines),
		),
	),
	styling.NewSetting(
		`^ {8}\t {12}\tDiff:\s*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetFaint),
			styling.NewSequence(sgr.ResetBoldAndFaint),
		),
	),
	styling.NewSetting(
		`^ {8}\t {12}\t--- Expected\s*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Blue)),
			styling.NewSequence(sgr.ResetForegroundColor),
		),
	),
	styling.NewSetting(
		`^ {8}\t {12}\t\+\+\+ Actual\s*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Red)),
			styling.NewSequence(sgr.ResetForegroundColor),
		),
	),
	styling.NewSetting(
		`^ {8}\t {12}\t@@ [-\+]?[\d]+,[-\+]?[\d]+ [-\+]?[\d]+,[-\+]?[\d]+ @@\s*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetFaint),
			styling.NewSequence(sgr.ResetBoldAndFaint),
		),
	),
	styling.NewSetting(
		`^ {8}\t {12}\t-.*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Blue)),
			styling.NewSequence(sgr.ResetForegroundColor),
		),
	),
	styling.NewSetting(
		`^ {8}\t {12}\t\+.*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Red)),
			styling.NewSequence(sgr.ResetForegroundColor),
		),
	),
	styling.NewSetting(
		`\s+time=(?:(?:"[^"]+")|(?:[0-9A-Za-z:\.]+)) level=DEBUG .*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Blue)),
			styling.NewSequence(sgr.ResetForegroundColor),
		),
	),
	styling.NewSetting(
		`\s+time=(?:(?:"[^"]+")|(?:[0-9A-Za-z:\.]+)) level=WARN .*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Yellow)),
			styling.NewSequence(sgr.ResetForegroundColor),
		),
	),
	styling.NewSetting(
		`\s+time=(?:(?:"[^"]+")|(?:[0-9A-Za-z:\.]+)) level=ERROR .*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Red)),
			styling.NewSequence(sgr.ResetForegroundColor),
		),
	),
	styling.NewSetting(
		`\s+[a-zA-Z0-9_\-\+]+=`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetFaint),
			styling.NewSequence(sgr.ResetBoldAndFaint),
		),
	),
	styling.NewSetting(
		`\s+time="[^"]+" level=DEBUG .*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.BrightBlack)),
			styling.NewSequence(sgr.ResetForegroundColor),
		),
	),
	styling.NewSetting(
		`\s+time="[^"]+" level=ERROR .*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Red)),
			styling.NewSequence(sgr.ResetForegroundColor),
		),
	),
	styling.NewSetting(
		`\s+time="[^"]+" level=WARN .*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Yellow)),
			styling.NewSequence(sgr.ResetForegroundColor),
		),
	),
	styling.NewSetting(
		`\s+[a-zA-Z0-9_\-\+]+=`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetFaint),
			styling.NewSequence(sgr.ResetBoldAndFaint),
		),
	),
}
