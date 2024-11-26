package knapsack_test

import (
	"dsa/dp/patterns/knapsack"
	"strconv"
	"testing"
)

var MincostTicketsTests = []struct {
	days, costs []int
	expected    int
}{
	{[]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}, 11},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}, 17},
}

func TestMincostTickets(t *testing.T) {
	for i, e := range MincostTicketsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := knapsack.MincostTicketsBottomUp(e.days, e.costs)
			if actual != e.expected {
				t.Fatalf("MincostTickets(%v, %v) = %v, want %v", e.days, e.costs, actual, e.expected)
			}
		})
	}
}
