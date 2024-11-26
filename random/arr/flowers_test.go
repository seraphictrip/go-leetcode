package arr_test

import (
	"dsa/random/arr"
	"slices"
	"strconv"
	"testing"
)

var CanPlaceFlowersTests = []struct {
	flowerbed []int
	n         int
	expected  bool
}{
	{[]int{1, 0, 0, 0, 1}, 1, true},
	{[]int{1, 0, 0, 0, 1}, 2, false},
}

func TestCanPlaceFlowers(t *testing.T) {
	for i, e := range CanPlaceFlowersTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			cp := slices.Clone(e.flowerbed)
			actual := arr.CanPlaceFlowers(e.flowerbed, e.n)
			if actual != e.expected {
				t.Fatalf("CanPlaceFlowers(%v, %v) = %v, want %v", cp, e.n, actual, e.expected)
			}
		})
	}
}
