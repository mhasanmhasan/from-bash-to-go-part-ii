package count_test

import (
	"count"
	"log"
	"strings"
	"testing"
)

func TestLinesCountsLines(t *testing.T) {
	t.Parallel()
	r := strings.NewReader("line1\nline2\nline1")
	want := map[string]int{
		"line1": 2,
		"line2": 1,
	}
	c, err := count.NewCounter(count.WithInput(r))
	if err != nil {
		log.Fatal(err)
	}
	got, err := c.Lines()
	if err != nil {
		log.Fatal(err)
	}
	if !mapsEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func mapsEqual(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k1, v1 := range m1 {
		if v2, exists := m2[k1]; !exists || v1 != v2 {
			return false
		}
	}
	return true
}

func TestCountCountsLinesFromArguments(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := count.NewCounter(count.WithInputFromArgs(args))
	if err != nil {
		log.Fatal(err)
	}
	want := map[string]int{
		"line1": 2,
		"line2": 1,
	}
	got, err := c.Lines()
	if err != nil {
		log.Fatal(err)
	}
	if !mapsEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCountCountsLinesFromArguments2(t *testing.T) {
	t.Parallel()
	tests := []struct {
		args []string
		want map[string]int
	}{
		{[]string{"testdata/three_lines.txt"}, map[string]int{"line1": 2, "line2": 1}},
		{[]string{}, map[string]int{}},
	}

	for _, test := range tests {
		c, err := count.NewCounter(count.WithInputFromArgs(test.args))
		if err != nil {
			log.Fatal(err)
		}
		got, err := c.Lines()
		if !mapsEqual(got, test.want) {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}

func TestWordsCountsWordsInInput(t *testing.T) {
	t.Parallel()
	r := strings.NewReader("one two three two three three")
	c, err := count.NewCounter(count.WithInput(r))
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	got, err := c.Words()
	if err != nil {
		t.Fatal(err)
	}
	if !mapsEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
