package subset_test

import (
	"dsa/backtracking/subset"
	"slices"
	"strconv"
	"testing"
)

var PowerSetTests = []struct {
	set      []int
	expected [][]int
}{
	// {[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
	// {[]int{1, 1, 2}, [][]int{{}, {1}, {2}, {1, 1}, {1, 2}, {1, 1, 2}}},
	{[]int{1, 2, 1}, [][]int{{}, {1}, {2}, {1, 1}, {1, 2}, {1, 1, 2}}},
}

func TestPowerSet(t *testing.T) {
	for i, e := range PowerSetTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := subset.PowerMultiSet(e.set)
			slices.SortFunc(actual, subsetSorter)
			slices.SortFunc(e.expected, subsetSorter)
			if len(actual) != len(e.expected) {
				t.Fatalf("PowerSet(%v) = %v, want %v", e.set, actual, e.expected)
			}
			for i := range actual {
				if !slices.Equal(actual[i], e.expected[i]) {
					t.Fatalf("PowerSet(%v) = %v, want %v", e.set, actual, e.expected)
				}
			}
		})
	}
}

func subsetSorter(a, b []int) int {
	m, n := len(a), len(b)
	if m == n {
		for i := 0; i < m; i++ {
			if a[i] != b[i] {
				return a[i] - b[i]
			}
		}
	}
	return m - n
}
