package dp_test

import (
	"dsa/advanced/dp"
	"strconv"
	"testing"
)

var RobTests = []struct {
	nums     []int
	expected int
}{
	{[]int{1, 2, 3, 1}, 4},
}

func TestRob(t *testing.T) {
	for i, e := range RobTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := dp.Rob(e.nums)
			if actual != e.expected {
				t.Fatalf("Rob(%v) = %v, want %v", e.nums, actual, e.expected)
			}
		})
	}
}
