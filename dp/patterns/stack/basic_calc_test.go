package stack_test

import (
	"dsa/dp/patterns/stack"
	"strconv"
	"testing"
)

var CalculateTests = []struct {
	expression string
	expected   int
}{
	{"3+2*2", 7},
	{" 3+5 / 2 ", 5},
	{"1-1+1", 1},
}

func TestCalculate(t *testing.T) {
	for i, e := range CalculateTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := stack.Calculate(e.expression)
			if actual != e.expected {
				t.Fatalf("Calculate(%v) = %v, want %v", e.expression, actual, e.expected)
			}
		})
	}
}
