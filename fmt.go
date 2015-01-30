package main

import (
	"fmt"

	"github.com/mgutz/ansi"
)

type colorPrinter struct {
	fn func(string) string
}

func colorFmt(style string) colorPrinter {
	return colorPrinter{
		fn: ansi.ColorFunc(style),
	}
}

func (p colorPrinter) Printf(format string, a ...interface{}) {
	fmt.Print(p.Sprintf(format, a...))
}

func (p colorPrinter) Sprintf(format string, a ...interface{}) string {
	return p.fn(fmt.Sprintf(format, a...))
}
