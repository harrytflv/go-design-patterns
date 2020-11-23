package bridge

import (
	"strings"
	"testing"
)

func TestPrinterAPI1(t *testing.T) {
	api1 := PrinterImpl1{}

	err := api1.PrintMessage("Hello")
	if err != nil {
		t.Fatal("?")
	}
}
func TestPrinterAPI2(t *testing.T) {
	api2 := PrinterImpl2{}

	expectedMessage := "Hello"

	if err := api2.PrintMessage(expectedMessage); err != nil {
		if !strings.Contains(err.Error(), "nil") {
			t.Fatal("?")
		}
	}

	testWriter := TestWriter{}
	api2 = PrinterImpl2{
		Writer: &testWriter,
	}

	if err := api2.PrintMessage(expectedMessage); err != nil {
		t.Fatalf("?")
	}

	if testWriter.Msg != expectedMessage {
		t.Fatalf("?")
	}
}

func TestNormalPrinter_Print(t *testing.T) {
	expectedMessage := "Hello io.Writer"

	normal := NormalPrinter{
		Msg:     expectedMessage,
		Printer: &PrinterImpl1{},
	}

	if err := normal.Print(); err != nil {
		t.Fatal("?")
	}

	testWriter := TestWriter{}
	normal = NormalPrinter{
		Msg: expectedMessage,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}

	if err := normal.Print(); err != nil {
		t.Fatal("?")
	}

	if testWriter.Msg != expectedMessage {
		t.Fatal("?")
	}
}

func TestPacktPrinter_Print(t *testing.T) {
	passedMessage := "Hello io.Writer"
	expectedMessage := "Message from Packt: Hello io.Writer"

	packt := PacktPrinter{
		Msg:     passedMessage,
		Printer: &PrinterImpl1{},
	}

	if err := packt.Print(); err != nil {
		t.Fatal("?")
	}

	testWriter := TestWriter{}
	packt = PacktPrinter{
		Msg: passedMessage,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}

	if err := packt.Print(); err != nil {
		t.Fatal("?")
	}

	if testWriter.Msg != expectedMessage {
		t.Fatalf("?")
	}
}
