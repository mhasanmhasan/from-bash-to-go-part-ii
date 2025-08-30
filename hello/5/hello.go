package hello

import (
	"io"
)

type Printer struct {
	Output io.Writer
}

func (p *Printer) Print() {
	s := "hello"
	p.Output.Write([]byte(s))
}
