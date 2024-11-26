package knapsack_test

import (
	"dsa/dp/patterns/knapsack"
	"strconv"
	"testing"
)

var KnapsackBruteForceTests = []struct {
	profit, weight []int
	cap            int
	expected       int
}{
	/*[
		[0 0 0 0 0 4 4 4 4]
		[0 0 4 4 4 4 4 8 8]
		[0 0 4 7 7 11 11 11 11]
		[0 1 4 7 8 11 12 12 12]
	]*/
	{
		[]int{4, 4, 7, 1},
		[]int{5, 2, 3, 1},
		8,
		12,
	},
}

func TestKnapsackBruteForce(t *testing.T) {
	for i, e := range KnapsackBruteForceTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := knapsack.KnapsackBottomUp(e.profit, e.weight, e.cap)
			if actual != e.expected {
				t.Fatalf("KnapsackBruteForce(%v, %v, %v) = %v, want %v", e.profit, e.weight, e.cap, actual, e.expected)
			}
		})
	}
}

var ConinChangeTests = []struct {
	amount   int
	coins    []int
	expected int
}{
	/*
		ORDER:
			terms i, amount requires:
				Take: i, amount - coins[i]
					- i does not change so amount depends on amount - coins[i]
					- smaller values before larger for amount

				Skip: i+1, amount
					- amount does not change so i depends on i
					-bigger values before smaller values for i

		[0,1,1,1,1,1]
		[0,1,2,2,3,0]
		[0,0,0,0,0,0]
		[0,0,0,0,0,0]
	*/
	{5, []int{1, 2, 5}, 4},
}

func TestConinChange(t *testing.T) {
	for i, e := range ConinChangeTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := knapsack.CoinChangeBottomUp(e.amount, e.coins)
			if actual != e.expected {
				t.Fatalf("CoinChange(%v, %v) = %v, want %v", e.amount, e.coins, actual, e.expected)
			}
		})
	}
}

var UnboundedKnapsackBruteForceTests = []struct {
	profit, weight []int
	cap            int
	expected       int
}{
	{
		[]int{4, 4, 7, 1},
		[]int{5, 2, 3, 1},
		8,
		18,
	},
}

func TestUnboundedKnapsackBruteForce(t *testing.T) {
	for i, e := range UnboundedKnapsackBruteForceTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := knapsack.UnboundedKnapsackBottomUp(e.profit, e.weight, e.cap)
			if actual != e.expected {
				t.Fatalf("UnboundKnapsack(%v, %v, %v) = %v, want %v", e.profit, e.weight, e.cap, actual, e.expected)
			}
		})
	}
}

var CoinChange1Tests = []struct {
	coins    []int
	amount   int
	expected int
}{
	{[]int{1, 2, 5}, 11, 3},
	{[]int{2}, 3, -1},
	{[]int{1}, 0, 0},
}

func TestCoinChange1(t *testing.T) {
	for i, e := range CoinChange1Tests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := knapsack.CoinChange1BottomUp(e.coins, e.amount)
			if actual != e.expected {
				t.Fatalf("CoinChange1(%v, %v) = %v, want %v", e.coins, e.amount, actual, e.expected)
			}
		})
	}
}
