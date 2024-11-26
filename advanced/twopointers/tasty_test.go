package twopointers_test

import (
	"dsa/advanced/twopointers"
	"strconv"
	"testing"
)

var TastyTests = []struct {
	price       []int
	k, expected int
}{
	{[]int{13, 5, 1, 8, 21, 2}, 3, 8},
	{[]int{1, 3, 1}, 2, 2},
	{[]int{7, 7, 7, 7}, 2, 0},
}

func TestTasty(t *testing.T) {
	for i, e := range TastyTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := twopointers.MaximumTastiness(e.price, e.k)
			if actual != e.expected {
				t.Fatalf("MaximumTastiness(%v, %v) = %v, want %v", e.price, e.k, actual, e.expected)
			}
		})
	}
}

var TastinessTests = []struct {
	c1, c2, expected int
}{
	{},
	{10, 5, 5},
	{5, 10, 5},
}

func TestTastiness(t *testing.T) {
	for i, e := range TastinessTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := twopointers.Tastiness(e.c1, e.c2)
			if actual != e.expected {
				t.Fatalf("Tastiness(%v, %v) = %v, want %v", e.c1, e.c2, actual, e.expected)
			}
		})
	}
}

var ChooseTests = []struct {
	n, k, expected int
}{
	{},
}

func TestChoose(t *testing.T) {
	for i, e := range ChooseTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := twopointers.Choose(e.n, e.k)
			if actual != e.expected {
				t.Fatalf("Choose(%v, %v) = %v, want %v", e.n, e.k, actual, e.expected)
			}
		})
	}
}

var BinarySearchTests = []struct {
	arr      []int
	target   int
	expected int
}{
	// {[]int{1}, 1, 0},
	// {[]int{1}, 2, -1},
	// {[]int{1}, -11, -1},
	// {[]int{1, 2}, 1, 0},

	{[]int{1, 2}, 2, 1},
}

func TestBinarySearch(t *testing.T) {
	for i, e := range BinarySearchTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := twopointers.BinarySearch(e.arr, e.target)
			if actual != e.expected {
				t.Fatalf("BinarySearch(%v, %v) = %v, want %v", e.arr, e.target, actual, e.expected)
			}
		})
	}
}
