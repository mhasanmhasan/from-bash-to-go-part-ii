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

func NewCounter(opts ...option) *counter {
	c := &counter{input: os.Stdin}
	for _, opt := range opts {
		opt(c) // NOTE: we ignore the error for now
	}
	return c
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

func (c *counter) Lines() (map[string]int, error) {
	counts := make(map[string]int)
	input := bufio.NewScanner(c.input)
	for input.Scan() {
		counts[input.Text()]++
	}
	return counts, input.Err()
}
