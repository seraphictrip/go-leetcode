package slidingwindow_test

import (
	"dsa/random/slidingwindow"
	"strconv"
	"testing"
)

var LongestOnesTests = []struct {
	nums        []int
	k, expected int
}{
	{

		[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6,
	},
	{
		[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1}, 3, 10,
	},
}

func TestLongestOnes(t *testing.T) {
	for i, e := range LongestOnesTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := slidingwindow.LongestOnes1(e.nums, e.k)
			if actual != e.expected {
				t.Fatalf("LongestOnes(%v, %v) = %v, want %v", e.nums, e.k, actual, e.expected)
			}
		})
	}
}
