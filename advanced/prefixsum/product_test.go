package prefixsum_test

import (
	"dsa/advanced/prefixsum"
	"slices"
	"strconv"
	"testing"
)

var ProductTests = []struct {
	nums, expected []int
}{
	{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
}

func TestProduct(t *testing.T) {
	for i, e := range ProductTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := prefixsum.ProductExceptSelfBruteForce(e.nums)
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("ProductExceptSelf(%v) = %v, want %v", e.nums, actual, e.expected)
			}
		})
	}
}
