package dp_test

import (
	"dsa/advanced/dp"
	"strconv"
	"testing"
)

var FibTests = []struct {
	n, expected int
}{
	// 0
	{0, 1},
	{1, 1},
	{2, 2},
	{3, 3},
	{4, 5},
	//				5
	//		 	  /	  		\
	// 		    4 			3
	//		/    \
	{5, 8},
	{6, 13},
	{7, 21},
	{8, 34},
	{9, 55},
	{10, 89},
	{11, 144},
	{12, 233},
	{13, 377},
	{14, 610},
	{42, 433494437},
}

func TestFib(t *testing.T) {
	for i, e := range FibTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := dp.FibDP(e.n)
			if actual != e.expected {
				t.Fatalf("Fib(%v) = %v, want %v", e.n, actual, e.expected)
			}
		})
	}
}

var fibs = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
