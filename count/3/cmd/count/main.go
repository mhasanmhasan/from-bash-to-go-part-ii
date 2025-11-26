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
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	
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
	select {
	case <-sig:
		println("\nProgram interrupted by user pressing Ctrl+C hence exiting gracefully!")
		os.Exit(0)
	default:
	}
}
