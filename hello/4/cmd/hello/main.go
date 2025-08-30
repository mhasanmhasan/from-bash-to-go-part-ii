package main

import (
	"hello"
	"os"
)

func main() {
	hello.Print() // defaults to os.Stdout

	hello.Output = os.Stderr
	hello.Print()
}
