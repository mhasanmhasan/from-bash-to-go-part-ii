package hello_test

import (
	"bytes"
	"hello"
	"testing"
)

func TestPrintToPrintsHelloToWriter(t *testing.T) {
	hello.Output = new(bytes.Buffer)
	hello.Print()
	want := "hello"
	got := hello.Output.(*bytes.Buffer).String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
