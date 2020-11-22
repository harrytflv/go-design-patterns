package adapter

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hello World!"

	adapter := PrinterAdapter{OldPrinter: &MyLegacyPrinter{}, Msg: msg}
	returnedMsg := adapter.PrintStored()
	if returnedMsg != "Legacy printer: Adapter: Hello World!\n" {
		t.Error("?")
	}

	adapter = PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnedMsg = adapter.PrintStored()
	if returnedMsg != "Hello World!" {
		t.Error("?")
	}
}
