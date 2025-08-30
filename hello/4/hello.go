package hello

import (
	"io"
	"os"
)

var Output io.Writer = os.Stdout

func Print() {
	s := "hello"
	Output.Write([]byte(s))
}
