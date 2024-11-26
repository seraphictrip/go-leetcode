package prefixsum_test

import (
	"dsa/patterns/prefixsum"
	"slices"
	"strconv"
	"testing"
)

var FindSubArraySumTests = []struct {
	input          []int
	i, j, expected int
}{
	{[]int{0}, 0, 0, 0},
	{[]int{1}, 0, 0, 1},
	{[]int{1, 2}, 0, 0, 1},
	{[]int{1, 2}, 0, 1, 3},
	{[]int{1, 2, 3}, 1, 1, 2},
	{[]int{1, 2, 3}, 1, 2, 5},
	{[]int{0, 1, 2, 3, 4, 5}, 1, 2, 3},
	{Range(0, 10, 1), 0, 9, 45},
	{Range(0, 100, 1), 0, 99, 4950},
	{Range(0, 1000, 1), 0, 999, 499500},
	{Range(1, 10, 1), 0, 9, 55},
	{Range(1, 100, 1), 0, 99, 5050},
	{Range(1, 1000, 1), 0, 999, 500500},
	{Range(1, 10000, 1), 0, 9999, 50005000},
	{Range(1, 100000, 1), 0, 99999, 5000050000},
	{Range(1, 1000000, 1), 0, 999999, 500000500000},
	{Range(1, 10000000, 1), 0, 9999999, 50000005000000},
	{Range(1, 100000000, 1), 0, 99999999, 5000000050000000},
}

func Range(start, count, inc int) []int {
	result := make([]int, count)
	for i := 0; i < count; i++ {
		result[i] = start + (i * inc)
	}
	return result
}

func TestFindSubArraySum(t *testing.T) {
	for i, e := range FindSubArraySumTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := prefixsum.FindSubArraySum(e.input, e.i, e.j)
			if actual != e.expected {
				t.Fatalf("FindSubArray(%v, %d, %d) = %d, want %d", "[...]", e.i, e.j, actual, e.expected)
			}
		})
	}
}

func TestSumFromPrefix(t *testing.T) {
	for i, e := range FindSubArraySumTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			prefixes := prefixsum.PrefixSum(e.input)
			actual := prefixsum.SumFromPrefix(prefixes, e.i, e.j)
			if actual != e.expected {
				t.Fatalf("SumFromPrefix(%v, %d, %d) = %d, want %d", "[...]", e.i, e.j, actual, e.expected)
			}
		})
	}
}

var PrefixSumTests = []struct {
	input []int
	last  int
}{
	{},
	{[]int{1, 2, 3, 4}, 10},
	{Range(0, 10, 1), 450},
	{Range(0, 100, 1), 4950},
	{Range(1, 100, 1), 5050},
	{Range(1, 1000, 1), 500500},
	{Range(1, 10000, 1), 50005000},
	{Range(1, 100000, 1), 5000050000},
	{Range(1, 1000000, 1), 500000500000},
	{Range(1, 10000000, 1), 50000005000000},
	{Range(1, 100000000, 1), 5000000050000000},
}

func TestPrefixSum(t *testing.T) {
	for i, e := range PrefixSumTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ps := prefixsum.PrefixSum2(e.input)
			if len(ps) > 0 && ps[len(ps)-1] != e.last {
				t.Fatalf("...%v", ps[len(ps)-9:])
			}
		})
	}
}

var PostfixSumTests = []struct {
	input, expected []int
}{
	{},
	{[]int{0}, []int{0}},
	{[]int{100}, []int{100}},
	{[]int{1, 2, 3}, []int{6, 5, 3}},
	{[]int{1, 2, 3, 4, 5}, []int{15, 14, 12, 9, 5}},
	{
		Range(1, 3, 1),
		reverse(prefixsum.PrefixSum(reverse(Range(1, 3, 1)))),
	},
}

func reverse(input []int) []int {
	slices.Reverse(input)
	return input
}

func TestPostfixSum(t *testing.T) {
	for i, e := range PostfixSumTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := prefixsum.PostfixSum2(e.input)
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("PostfixSum(%v) = %v, want %v", e.input, actual, e.expected)
			}
		})
	}
}
