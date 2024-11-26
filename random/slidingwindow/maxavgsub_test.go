package slidingwindow_test

import (
	"dsa/random/slidingwindow"
	"strconv"
	"testing"
)

var FindMaxAverageTests = []struct {
	nums     []int
	k        int
	expected float64
}{
	{
		[]int{1, 12, -5, -6, 50, 3}, 4, 12.75,
	},
	{
		[]int{5}, 1, 5.0,
	},
}

func TestFindMaxAverage(t *testing.T) {
	for i, e := range FindMaxAverageTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := slidingwindow.FindMaxAverage(e.nums, e.k)
			if actual != e.expected {
				t.Fatalf("FindMaxAverage(%v, %v) = %0.5f, want %0.5f", e.nums, e.k, actual, e.expected)
			}
		})
	}
}
