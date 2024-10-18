package main

import (
	"github.com/pamburus/go-ansi-esc/sgr"
	"github.com/pamburus/go-mod/build/tools/internal/styling"
)

var styleSettings = []styling.Setting{
	styling.NewSetting(
		`^\s*--- PASS: .* \(\d+\.\d+s\)$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Green)),
			styling.NewSequence(sgr.ResetAll),
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
			styling.NewSequence(sgr.SetForegroundColor(sgr.White), sgr.SetBackgroundColor(sgr.Green), sgr.SetBold)+"[",
			"]"+styling.NewSequence(sgr.ResetAll),
		),
	),
	styling.NewSetting(
		`^FAIL$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.White), sgr.SetBackgroundColor(sgr.Red), sgr.SetBold)+"[",
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
			styling.NewSequence(),
			styling.NewSequence(),
		),
	),
	styling.NewSetting(
		`\bactual\s*: .*`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Red), sgr.SetUnderlined),
			styling.NewSequence(sgr.ResetAll),
		),
	),
}
