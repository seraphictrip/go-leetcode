package knapsack_test

import (
	"dsa/dp/patterns/knapsack"
	"strconv"
	"testing"
)

var FindMaxFormTests = []struct {
	strs     []string
	m        int
	n        int
	expected int
}{
	// {[]string{"10", "0001", "111001", "1", "0"}, 5, 3, 4},
	// {[]string{"10", "0", "1"}, 1, 1, 2},
	// {[]string{"10", "0001", "111001", "1", "0"}, 3, 4, 3},
	// {[]string{"0", "11", "1000", "01", "0", "101", "1", "1", "1", "0", "0", "0", "0", "1", "0", "0110101", "0", "11", "01", "00", "01111", "0011", "1", "1000", "0", "11101", "1", "0", "10", "0111"}, 9, 90, 17},
	{[]string{"011", "1", "11", "0", "010", "1", "10", "1", "1", "0", "0", "0", "01111", "011", "11", "00", "11", "10", "1", "0", "0", "0", "0", "101", "001110", "1", "0", "1", "0", "0", "10", "00100", "0", "10", "1", "1", "1", "011", "11", "11", "10", "10", "0000", "01", "1", "10", "0"}, 44, 39, 45},
}

func TestFindMaxForm(t *testing.T) {
	for i, e := range FindMaxFormTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := knapsack.FindMaxFormBottomUp(e.strs, e.m, e.n)
			if actual != e.expected {
				t.Fatalf("FindMaxForm(%v, %v, %v) = %v, want %v", e.strs, e.m, e.n, actual, e.expected)
			}
		})
	}
}
