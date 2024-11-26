package twopointers_test

import (
	"dsa/advanced/twopointers"
	"slices"
	"strconv"
	"testing"
)

var IsPalindromeTests = []struct {
	arr      []int
	expected bool
}{
	{[]int{1, 2, 7, 7, 2, 1}, true},
	{[]int{1, 2, 7, 3, 7, 2, 1}, true},
}

func TestIsPalindrome(t *testing.T) {
	for i, e := range IsPalindromeTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := twopointers.IsPalindrome2(e.arr)

			if actual != e.expected {
				t.Fatalf("IsPalindrome(%v) = %v, want %v", e.arr, actual, e.expected)
			}
		})
	}
}

var IsCleansedPalindromeTests = []struct {
	s        string
	expected bool
}{
	{"A man, a plan, a canal: Panama", true},
}

func TestIsCleansedPalindrome(t *testing.T) {
	for i, e := range IsCleansedPalindromeTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := twopointers.IsPalindrome3(e.s)
			if actual != e.expected {
				t.Fatalf("IsCleansedPalindrome(%v) = %v, want %v", e.s, actual, e.expected)
			}
		})
	}
}

var RemoveDuplicatesTests = []struct {
	arr, expected []int
}{
	// [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
	{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
}

func TestRemoveDuplicates(t *testing.T) {
	for i, e := range RemoveDuplicatesTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			cp := slices.Clone(e.arr)
			actual := twopointers.RemoveDuplicatesConcept(e.arr)
			if actual != len(e.expected) {
				t.Fatalf("RemoveDuplicates(%v) = %v, want %v", cp, actual, e.expected)
			}
		})
	}
}
