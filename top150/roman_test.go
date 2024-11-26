package arr_test

import (
	arr "dsa/top150"
	"strconv"
	"testing"
)

var RomanTests = []struct {
	roman    string
	expected int
}{
	{"III", 3},
	{"I", 1},
	{"V", 5},
	{"IV", 4},
	{"VI", 6},
}

func TestRoman(t *testing.T) {
	for i, e := range RomanTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := arr.RomanToInt(e.roman)
			if actual != e.expected {
				t.Fatalf("RomanToInt(%q) = %v, want %v", e.roman, actual, e.expected)
			}
		})
	}
}
