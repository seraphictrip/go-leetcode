package groupanagrams_test

import (
	"slices"
	"strconv"
	"testing"
)

var SortStringTests = []struct {
	input, expected string
}{
	{"ate", "aet"},
	{"eat", "aet"},
	{"tea", "aet"},
}

func TestSortString(t *testing.T) {
	for i, e := range SortStringTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := sortString(e.input)
			if actual != e.expected {
				t.Fatalf("sortString(%v) = %v, want %v", e.input, actual, e.expected)
			}
		})
	}
}

func sortString(s string) string {
	runes := []rune(s)
	slices.Sort(runes)
	return string(runes)
}
