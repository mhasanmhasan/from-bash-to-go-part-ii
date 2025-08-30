package hello

import "io"

func PrintTo(w io.Writer) {
	s := "hello"
	w.Write([]byte(s))
}
