package cycle

import (
	"strconv"
	"testing"
)

var HasCycleTests = []struct {
	list     *ListNode
	expected bool
}{
	{},
	{FromArray([]int{1, 2, 3}), false},
	{FromArrayCycleAt([]int{1, 2, 3, 4, 5}, 0), true},
	{FromArrayCycleAt(Range(0, 100, 1), 78), true},
}

func Range(start, end, inc int) []int {
	size := end - start
	rang := make([]int, size)
	for i := 0; i < size; i++ {
		rang[i] = start + (inc * i)
	}
	return rang
}
func TestHasCycle(t *testing.T) {
	for i, e := range HasCycleTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := HasCycle(e.list)
			Print(e.list)
			if actual != e.expected {
				t.Fatalf("HasCycle(%v) = %v, want %v", e.list, actual, e.expected)
			}
		})
	}
}
