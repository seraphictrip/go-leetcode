package arr_test

import (
	"dsa/random/arr"
	"slices"
	"strconv"
	"testing"
)

var ProductExceptSelfTests = []struct {
	nums     []int
	expected []int
}{
	{
		[]int{1, 2, 3, 4},
		[]int{24, 12, 8, 6},
	},
}

func TestProductExceptSelf(t *testing.T) {
	for i, e := range ProductExceptSelfTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := arr.ProductExceptSelf(e.nums)
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("ProductExceptSelf(%v) = %v, want %v", e.nums, actual, e.expected)
			}
		})
	}
}
