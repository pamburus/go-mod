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
			styling.NewSequence(sgr.ResetAll),
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
			styling.NewSequence(sgr.ResetAll),
		),
	),
	styling.NewSetting(
		`^(ok|\?)\s+.*`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetFaint),
			styling.NewSequence(sgr.ResetAll),
		),
	),
	styling.NewSetting(
		`^=== RUN\s.*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetFaint),
			styling.NewSequence(sgr.ResetAll),
		),
	),
	styling.NewSetting(
		`[^\s]+\.go(:\d+)?`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Cyan)),
			styling.NewSequence(sgr.ResetAll),
		),
	),
	styling.NewSetting(
		`\s*Error(?: Trace)?:.*`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Red)),
			styling.NewSequence(sgr.ResetAll),
		),
	),
	styling.NewSetting(
		`\bexpected\s*: .*`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Blue)),
			styling.NewSequence(sgr.ResetAll),
		),
	),
	styling.NewSetting(
		`\bactual\s*: .*`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Red), sgr.SetUnderlined),
			styling.NewSequence(sgr.ResetAll),
		),
	),
	styling.NewSetting(
		`^ {8}\t {12}\tDiff:\s*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetFaint),
			styling.NewSequence(sgr.ResetAll),
		),
	),
	styling.NewSetting(
		`^ {8}\t {12}\t--- Expected\s*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Blue)),
			styling.NewSequence(sgr.ResetAll),
		),
	),
	styling.NewSetting(
		`^ {8}\t {12}\t\+\+\+ Actual\s*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Red)),
			styling.NewSequence(sgr.ResetAll),
		),
	),
	styling.NewSetting(
		`^ {8}\t {12}\t@@ [-\+]?[\d]+,[-\+]?[\d]+ [-\+]?[\d]+,[-\+]?[\d]+ @@\s*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetFaint),
			styling.NewSequence(sgr.ResetAll),
		),
	),
	styling.NewSetting(
		`^ {8}\t {12}\t- .*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Blue)),
			styling.NewSequence(sgr.ResetAll),
		),
	),
	styling.NewSetting(
		`^ {8}\t {12}\t\+ .*$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Red)),
			styling.NewSequence(sgr.ResetAll),
		),
	),
}
