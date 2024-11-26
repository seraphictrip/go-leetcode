package arrays_test

import (
	"dsa/arrays"
	"strconv"
	"testing"
)

var ConvertTests = []struct {
	s        string
	numRows  int
	expected string
}{
	// [0, 4, 8, 12] f(x) = 0+4
	// [1, 3, 5, 7, 9, 11, 13] f(x) = 1+2
	// [2, 6, 10] f(x) = 2+4
	{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
	// [0, 6, 12] f(x) = 0+6
	// [1, 5, 7, 11, 13]
	// [2, 4 ,8, 10]
	// [3, 9]
	{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", 8, ""},
}

func TestConvert(t *testing.T) {
	for i, e := range ConvertTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := arrays.Convert(e.s, e.numRows)
			if actual != e.expected {
				t.Fatalf("Convert(%q, %v) = %q, want %q", e.s, e.numRows, actual, e.expected)
			}
		})
	}
}
