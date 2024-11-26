package backtracking_test

import (
	"dsa/backtracking"
	"slices"
	"strconv"
	"testing"
)

var BinaryWatchTests = []struct {
	turnedOn int
	expected []string
}{
	{1, []string{}},
}

func TestBinaryWatch(t *testing.T) {
	for i, e := range BinaryWatchTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := backtracking.ReadBinaryWatch(e.turnedOn)
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("ReadBinaryWatch(%v) = %v, want %v", e.turnedOn, actual, e.expected)
			}
		})
	}
}
