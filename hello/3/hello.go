package hello

import (
	"io"
	"os"
)

func PrintTo(w io.Writer) {
	if w == nil {
		w = os.Stdout
	}
	s := "hello"
	w.Write([]byte(s))
}
