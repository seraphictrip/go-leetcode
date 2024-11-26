package dp_test

import (
	"dsa/dp"
	"strconv"
	"testing"
)

var CoinChangeTests = []struct {
	amount   int
	coins    []int
	expected int
}{
	// {0, []int{1}, 1},
	//		change(5, [1,2,5])
	//	change(4,[1,2,4])	change(5, [2,5])
	{5, []int{1, 2, 5}, 4},
	// {3, []int{2}, 0},
	// {10, []int{10}, 1},
	{500, []int{3, 5, 7, 8, 9, 10, 11}, 35502874},
}

func TestCoinChange(t *testing.T) {
	for i, e := range CoinChangeTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := dp.Change(e.amount, e.coins)
			if actual != e.expected {
				t.Fatalf("Change(%v, %v) = %v, want %v", e.amount, e.coins, actual, e.expected)
			}
		})
	}
}
