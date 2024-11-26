package backtracking_test

import (
	bt "dsa/backtracking"
	"fmt"
	"slices"
	"strconv"
	"testing"
)

var PowerSetTests = []struct {
	set         []int
	cardinality int
	ordered     [][]int
}{
	{[]int{1, 2, 3}, 8, [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
	{[]int{2, 7, 9}, 8, [][]int{{}, {2}, {7}, {9}, {2, 7}, {2, 9}, {7, 9}, {2, 7, 9}}},
	{[]int{1, 2, 3, 4}, 16, [][]int{
		{},
		{1}, {2}, {3}, {4},
		{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4},
		{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4},
		{1, 2, 3, 4},
	}},
}

func TestPowerSet(t *testing.T) {
	for i, e := range PowerSetTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ps := bt.PowerSet(e.set)
			sortps(ps)
			fmt.Println(ps)
			if len(ps) != bt.PowerSetCardinality(e.set) {
				t.Fatalf("PowerSetCardinality(%v) = %v, want %v |P(A)| = 2^n", e.set, len(ps), bt.PowerSetCardinality(e.set))
			}
			AssertSamePowerSet(t, e.set, ps, e.ordered)

		})
	}
}

// Sort a power set
// first by cardinality, then lexigraphically
func sortps(ps [][]int) {
	slices.SortStableFunc(ps, func(a, b []int) int {
		n, m := len(a), len(b)
		if n < m {
			return -1
		}
		if n > m {
			return 1
		}
		// are the same len, find first char that doesn't match
		// and use that to sort
		for i := 0; i < n; i++ {
			if a[i] != b[i] {
				return a[i] - b[i]
			}
		}
		return 0
	})
}

func AssertSamePowerSet(t *testing.T, set []int, actual, expected [][]int) {
	t.Helper()
	if len(actual) != len(expected) {
		t.Fatalf("PowerSet(%v) = %v, want %v", set, actual, expected)
	}
	for i := 0; i < len(actual); i++ {
		if !slices.Equal(actual[i], expected[i]) {
			t.Fatalf("PowerSet(%v) = %v, want %v: see row %v", set, actual, expected, i)
		}
	}
}
