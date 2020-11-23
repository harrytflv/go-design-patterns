package bridge

import (
	"errors"
	"fmt"
	"io"
)

type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Println(msg)
	return nil
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (p *PrinterImpl2) PrintMessage(msg string) error {
	if p.Writer == nil {
		return errors.New("nil")
	}

	fmt.Fprintf(p.Writer, "%s", msg)
	return nil
}

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (int, error) {
	n := len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	return 0, errors.New("Empty")
}

type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *NormalPrinter) Print() error {
	return c.Printer.PrintMessage(c.Msg)
}

type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *PacktPrinter) Print() error {
	return c.Printer.PrintMessage("Message from Packt: " + c.Msg)
}
