package main

import (
	"github.com/pamburus/go-ansi-esc/sgr"
	"github.com/pamburus/go-mod/build/tools/internal/styling"
)

var styleSettings = []styling.Setting{
	styling.NewSetting(
		`\b100\.0%$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Blue)),
			styling.NewSequence(sgr.ResetAll),
		),
	),
	styling.NewSetting(
		`\b[0-9]{2}\.[0-9]%$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Yellow)),
			styling.NewSequence(sgr.ResetAll),
		),
	),
	styling.NewSetting(
		`\b[0-9]\.[0-9]%$`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.BrightRed)),
			styling.NewSequence(sgr.ResetAll),
		),
	),
	styling.NewSetting(
		`^[^\s]+\.go(:\d+)?:?`,
		styling.NewStyle(
			styling.NewSequence(sgr.SetForegroundColor(sgr.Cyan)),
			styling.NewSequence(sgr.ResetAll),
		),
	),
}
