package recurrence

import (
	"strconv"
	"testing"
)

var PathsTests = []struct {
	n, m, expected int
}{
	{},
	{1, 1, 1},
	{2, 2, 2},
	{2, 4, 4},
	// [][][]
	// [][][]
	// [][][]
	{3, 3, 6},
}

func TestPaths(t *testing.T) {
	for i, e := range PathsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := Paths(e.n, e.m)
			if actual != e.expected {
				t.Fatalf("Paths(%v, %v) = %v, want %v", e.n, e.m, actual, e.expected)
			}
		})
	}
}
