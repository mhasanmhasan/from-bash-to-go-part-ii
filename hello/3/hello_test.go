package hello_test

import (
	"bytes"
	"hello"
	"testing"
)

func TestPrintToPrintsHelloToWriter(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	hello.PrintTo(buf)
	want := "hello"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
