package main

import (
	"count"
	"flag"
	"fmt"
	"log"
)

const usage = `Counts words (or lines) from stdin (or files).

Usage: count [-lines] [file...]`

func main() {
	flag.Usage = func() {
		fmt.Println(usage)
		flag.PrintDefaults()
	}
	lines := flag.Bool("lines", false, "count lines, not words")
	flag.Parse()

	c, err := count.NewCounter(
		count.WithInputFromArgs(flag.Args()),
	)
	if err != nil {
		log.Fatal(err)
	}

	var counts map[string]int
	if *lines {
		counts, err = c.Lines()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		counts, err = c.Words()
		if err != nil {
			log.Fatal(err)
		}
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
