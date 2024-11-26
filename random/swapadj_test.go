package random_test

import (
	"dsa/random"
	"strconv"
	"testing"
)

var CanTransformTests = []struct {
	start, end string
	expected   bool
}{
	// {"X", "L", false},
	{"RXXLRXRXL", "XRLXXRRLX", true},
}

func TestCanTransform(t *testing.T) {
	for i, e := range CanTransformTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := random.CanTransform(e.start, e.end)
			if actual != e.expected {
				t.Fatalf("CanTransform(%q, %q) = %v, want %v", e.start, e.end, actual, e.expected)
			}
		})
	}
}
