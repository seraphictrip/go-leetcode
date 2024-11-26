package stack_test

import (
	"dsa/dp/patterns/stack"
	"strconv"
	"testing"
)

var SimplifyPathTests = []struct {
	path, expected string
}{
	{"/home//foo/", "/home/foo"},
	{"/home/user/Documents/../Pictures", "/home/user/Pictures"},
	{"/../", "/"},
}

func TestSimplifyPath(t *testing.T) {
	for i, e := range SimplifyPathTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := stack.SimplifyPath(e.path)
			if actual != e.expected {
				t.Fatalf("SimplifyPath(%v) = %v, want %v", e.path, actual, e.expected)
			}
		})
	}
}
