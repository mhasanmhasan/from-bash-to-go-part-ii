package main

import (
	"count"
	"fmt"
	"log"
	"os"
)

func main() {
	c, err := count.NewCounter(count.WithInputFromArgs(os.Args[1:]))
	if err != nil {
		log.Fatal(err)
	}
	counts, err := c.Lines()
	if err != nil {
		log.Fatal(err)
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
