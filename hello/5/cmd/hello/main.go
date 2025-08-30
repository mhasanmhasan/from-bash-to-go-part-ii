package main

import (
	"hello"
	"os"
)

func main() {
	p := hello.Printer{Output: os.Stdout}
	p.Print()
}
