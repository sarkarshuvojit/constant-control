package util

import "github.com/sarkarshuvojit/pprinter/pprinter"

var (
	Printer pprinter.Pprinter = *pprinter.WithTheme(&pprinter.AyuTheme)
)
