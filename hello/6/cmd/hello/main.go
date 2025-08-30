package main

import (
	"hello"
	"os"
)

func main() {
	// hello.NewPrinter().Print()
	p := hello.NewPrinter()
	p.Output = os.Stderr
	p.Print()
}
