package guessnumber_test

import (
	"dsa/binarysearch/guessnumber"
	"math"
	"strconv"
	"testing"
)

var GuessNumberTests = []struct {
	// 1 <= n <= math.Pow(2, 31) - 1
	n int
	// 1 <= pick <= n
	pick int
}{
	// {}, // not valid, undefined as range is 1 <= n <=math.Pow(2, 31) - 1
	{1, 1},
	{10, 7},
	{math.MaxInt16, 8},
	{math.MaxInt16, 15734},
	{math.MaxInt16, math.MaxInt16},
	{math.MaxInt16, 0},
}

func TestGuessNumber(t *testing.T) {
	for i, e := range GuessNumberTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// fmt.Println(math.MaxUint32>>1, math.MaxUint32/2, math.Pow(2, 31)-1)
			guessFn := guessnumber.MakeGuess(e.pick)
			actual := guessnumber.GuessNumber(e.n, guessFn)
			if actual != e.pick {
				t.Fatalf("GuessNumber(%v) = %v, want %v", e.n, actual, e.pick)
			}
		})
	}
}

var GetMoneyAmountTests = []struct {
	n        int
	expected int
}{
	// {10, 7},
	// {1, 0},
	// {2, 1},
	// {3, 2},
	//			(2)
	//		(1)		(4)
	//					(5)
	// guess 1: 2, minCost = 2 + min(left, right)
	{4, 5},
}

func TestGetMoneyAmount(t *testing.T) {
	for i, e := range GetMoneyAmountTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := guessnumber.GetMoneyAmount(e.n)
			if actual != e.expected {
				t.Fatalf("GetMoneyAmount(%v) = %v, want %v", e.n, actual, e.expected)
			}
		})
	}
}
