package main

import (
	"count"
	"fmt"
	"log"
	"os"
)

func main() {
	var counts map[string]int
	var err error

	if len(os.Args) > 1 {
		// Input from file.
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		c := count.NewCounter(count.WithInput(file))
		counts, err = c.Lines()
	} else {
		// Input from stdin.
		c := count.NewCounter()
		counts, err = c.Lines()
	}

	if err != nil {
		log.Fatal(err)
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
