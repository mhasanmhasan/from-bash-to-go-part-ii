package hello

import (
	"io"
	"os"
)

type Printer struct {
	Output io.Writer
}

func NewPrinter() *Printer {
	return &Printer{
		Output: os.Stdout,
	}
}

func (p *Printer) Print() {
	s := "hello"
	p.Output.Write([]byte(s))
}
