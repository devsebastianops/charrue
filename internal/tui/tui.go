package tui

import (
	"github.com/devsebastianops/charrue/internal/common"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func Header() {
	pterm.Println()

	style := pterm.NewStyle(pterm.FgGreen, pterm.Bold, pterm.Underscore)
	s, _ := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("CHARRUE", style),
	).Srender()
	pterm.DefaultCenter.Println(s)

	pterm.DefaultCenter.Println(common.ShortDescription)
	pterm.DefaultCenter.Println("Version: " + pterm.Green(common.Version))
}

func Info(msg string) {
	pterm.Info.Println(msg)
}

func Error(msg string) {
	pterm.Error.Println(msg)
}

func Success(msg string) {
	pterm.Success.Println(msg)
}

func EnableDebug() {
	pterm.EnableDebugMessages()
}

func Debug(msg string, context ...any) {
	if len(context) > 0 {
		msg = msg + " " + pterm.Sprintf("%+v", context)
	}

	pterm.Debug.Println(msg)
}
