package dp_test

import (
	"dsa/dp"
	"strconv"
	"testing"
)

var FibTests = []struct {
	n, expected int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
	{10, 55},
	{11, 89},
	{12, 144},
	{42, 267914296},
}

func TestFib(t *testing.T) {
	for i, e := range FibTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := dp.FibMemo(e.n)
			if actual != e.expected {
				t.Fatalf("Fib(%v) = %v, want %v", e.n, actual, e.expected)
			}
		})
	}
}
