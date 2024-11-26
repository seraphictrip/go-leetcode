package recurrence_test

import (
	"dsa/advanced/recurrence"
	"strconv"
	"testing"
)

var PartitionsTests = []struct {
	n, m     int
	expected int
}{
	{0, 0, 1},
	{0, 1, 1},
	{0, 2, 1},
	{1, 0, 0},
	{2, 0, 0},
	// [*]
	{1, 1, 1},
	// [*]+[*]
	{2, 1, 1},
	// [*]+[*]+[*]
	{3, 1, 1},
}

func TestPartitions(t *testing.T) {
	for i, e := range PartitionsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := recurrence.Partitions(e.n, e.m)
			if actual != e.expected {
				t.Fatalf("Partitions(%v) = %v, want %v", e.n, actual, e.expected)
			}
		})
	}
}
