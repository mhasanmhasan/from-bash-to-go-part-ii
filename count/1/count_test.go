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
	c := count.NewCounter(count.WithInput(r))
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
