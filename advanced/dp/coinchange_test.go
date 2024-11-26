package dp_test

import (
	"dsa/advanced/dp"
	"strconv"
	"testing"
)

var CoinChangeTests = []struct {
	coins    []int
	amount   int
	expected int
}{
	{[]int{1, 2, 5}, 11, 3},
	{[]int{2}, 3, -1},
	{[]int{1}, 0, 0},
}

func TestCoinChange(t *testing.T) {
	for i, e := range CoinChangeTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := dp.CoinChange(e.coins, e.amount)
			if actual != e.expected {
				t.Fatalf("CoinChange(%v, %v) = %v, want %v", e.coins, e.amount, actual, e.expected)
			}
		})
	}
}
