package arr_test

import (
	"dsa/random/arr"
	"strconv"
	"testing"
)

var IncreasingTripletTests = []struct {
	nums     []int
	expected bool
}{
}

func TestIncreasingTriplet(t *testing.T) {
	for i, e := range IncreasingTripletTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := arr.IncreasingTripletPrefix(e.nums)
			if actual != e.expected {
				t.Fatalf("IncreasingTriplet(%v) = %v, want %v", e.nums, actual, e.expected)
			}
		})
	}
}