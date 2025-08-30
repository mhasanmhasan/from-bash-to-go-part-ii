package count

import (
	"bufio"
	"errors"
	"io"
	"os"
)

type counter struct {
	input io.Reader
}

type option func(*counter) error

func NewCounter(opts ...option) (*counter, error) {
	c := &counter{input: os.Stdin}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func WithInputFromArgs(args []string) option {
	return func(c *counter) error {
		if len(args) < 1 {
			return nil
		}
		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		c.input = f
		// NOTE: We are not closing the f and we take only the first
		// argument. See count/3/count.go for how to fix both
		// shortcomings.
		return nil
	}
}

func (c *counter) Lines() (map[string]int, error) {
	counts := make(map[string]int)
	input := bufio.NewScanner(c.input)
	for input.Scan() {
		counts[input.Text()]++
	}
	return counts, input.Err()
}
