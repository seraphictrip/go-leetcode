package prefixsum_test

import (
	"dsa/advanced/prefixsum"
	"slices"
	"strconv"
	"testing"
)

var inclusiveTests = []struct {
	nums     []int
	expected []int
}{
	{[]int{1, 2}, []int{1, 3}},
	{[]int{1, 2, 3, 4, 5}, []int{1, 3, 6, 10, 15}},
	{[]int{1, 1, 2, 3}, []int{1, 2, 4, 7}},
}

func TestInclusive(t *testing.T) {
	for i, e := range inclusiveTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := prefixsum.PrefixSumInclusive(e.nums)
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("PrefixSumInclusive(%v) = %v, want %v", e.nums, actual, e.expected)
			}
		})
	}
}

var ExclusiveTests = []struct {
	nums, expected []int
}{
	{[]int{1, 2, 3, 4, 5}, []int{0, 1, 3, 6, 10}},
}

func TestExclusive(t *testing.T) {
	for i, e := range ExclusiveTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := prefixsum.PrefixSumExclusive(e.nums)
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("PrefixSumExclusive(%v) = %v, want %v", e.nums, actual, e.expected)
			}
		})
	}
}
