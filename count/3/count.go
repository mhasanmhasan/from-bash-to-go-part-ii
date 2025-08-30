package count

import (
	"bufio"
	"errors"
	"io"
	"os"
)

type counter struct {
	input io.Reader
	files []io.Reader
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
		c.files = make([]io.Reader, len(args))
		for i, path := range args {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			c.files[i] = f
		}
		c.input = io.MultiReader(c.files...)
		return nil
	}
}

func (c *counter) Lines() (map[string]int, error) {
	counts := make(map[string]int)
	input := bufio.NewScanner(c.input)
	for input.Scan() {
		counts[input.Text()]++
	}
	for _, f := range c.files {
		if closer, ok := f.(io.Closer); ok {
			closer.Close()
		}
	}
	return counts, input.Err()
}

func (c *counter) Words() (map[string]int, error) {
	counts := make(map[string]int)
	input := bufio.NewScanner(c.input)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
	for _, f := range c.files {
		if closer, ok := f.(io.Closer); ok {
			closer.Close()
		}
	}
	return counts, input.Err()
}
