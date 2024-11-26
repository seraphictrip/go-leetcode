package hashmap_test

import (
	"dsa/hashmap"
	"slices"
	"strconv"
	"testing"
)

var GetDigitsTests = []struct {
	n        int
	expected []int
}{
	{1, []int{1}},
	{10, []int{0, 1}},
	{11, []int{1, 1}},
}

func TestGetDigits(t *testing.T) {
	for i, e := range GetDigitsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := hashmap.GetDigits(e.n)
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("GetDigits(%v) = %v, want %v", e.n, actual, e.expected)
			}
		})
	}
}

var IsHappyTests = []struct {
	n        int
	expected bool
}{
	{19, true},
}

func TestIsHappy(t *testing.T) {
	for i, e := range IsHappyTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := hashmap.IsHappy(e.n)
			if actual != e.expected {
				t.Fatalf("IsHappy(%v) = %v, want %v", e.n, actual, e.expected)
			}
		})
	}
}
