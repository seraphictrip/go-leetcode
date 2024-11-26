package stack_test

import (
	"dsa/dp/patterns/stack"
	"strconv"
	"testing"
)

var DecodeStringTests = []struct {
	s, expected string
}{
	{"3[a]2[bc]", "aaabcbc"},
	/*
		2
		"a"
		3
		""
	*/
	{"3[a2[c]]", "accaccacc"},
}

func TestDecodeString(t *testing.T) {
	for i, e := range DecodeStringTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := stack.DecodeString(e.s)
			if actual != e.expected {
				t.Fatalf("DecodeString(%v) = %v, want %v", e.s, actual, e.expected)
			}
		})
	}
}

var EvalRPNTests = []struct {
	tokens   []string
	expected int
}{
	{
		[]string{"2", "1", "+", "3", "*"}, 9,
	},
	{[]string{"4", "13", "5", "/", "+"}, 6},
}

func TestEvalRPN(t *testing.T) {
	for i, e := range EvalRPNTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := stack.EvalRPN(e.tokens)
			if actual != e.expected {
				t.Fatalf("EvalRPN(%v) = %v, want %v", e.tokens, actual, e.expected)
			}
		})
	}
}
