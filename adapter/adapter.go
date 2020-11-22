package adapter

import (
	"fmt"
)

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct {
}

func (l *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy printer: %s\n", s)
	println(newMsg)
	return
}

type ModernPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = p.OldPrinter.Print(fmt.Sprintf("Adapter: %s", p.Msg))
	} else {
		newMsg = p.Msg
	}
	return
}
