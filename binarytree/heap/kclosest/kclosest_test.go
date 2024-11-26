package kclosest_test

import (
	"dsa/binarytree/heap/kclosest"
	"math"
	"slices"
	"strconv"
	"testing"
)

var EuclideanDistanceTests = []struct {
	x1, y1, x2, y2, expected float64
}{
	{},
	{1, 1, 0, 0, math.Sqrt(2)},
	{1, 0, 0, 1, math.Sqrt(2)},
	{2, 0, 0, 0, 2},
	{2, 2, 0, 0, math.Sqrt(8)},
	{1, 3, 0, 0, math.Sqrt(10)},
	{-2, 2, 0, 0, math.Sqrt(8)},
	{3, 3, 0, 0, math.Sqrt(18)},
	{5, -1, 0, 0, math.Sqrt(26)},
	{-2, 4, 0, 0, math.Sqrt(20)},
}

func TestEuclideanDistance(t *testing.T) {
	for i, e := range EuclideanDistanceTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := kclosest.EuclideanDistance(e.x1, e.y1, e.x2, e.y2)
			if actual != e.expected {
				t.Fatalf("EuclideanDistance(%v, %v, %v, %v) = %v, want %v", e.x1, e.y1, e.x2, e.y2, actual, e.expected)
			}
		})
	}
}

var kClosestTests = []struct {
	points   [][]int
	k        int
	expected [][]int
}{
	{[][]int{{1, 3}, {-2, 2}}, 1, [][]int{{-2, 2}}},
	{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2, [][]int{{3, 3}, {-2, 4}}},
}

func TestKClosest(t *testing.T) {
	for i, e := range kClosestTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := kclosest.KClosestHeap(e.points, e.k)
			for i := range actual {
				if !slices.Equal(actual[i], e.expected[i]) {
					t.Fatalf("KClosest(...) = %v, want %v", actual, e.expected)
				}
			}
		})
	}
}

func asc(a, b int) int {
	return a - b
}

func desc(a, b int) int {
	return b - a
}

var SortTests = []struct {
	input    []int
	fn       func(a, b int) int
	expected []int
}{
	{
		[]int{1, 2, 3, 4, 5},
		asc,
		[]int{1, 2, 3, 4, 5},
	},
	{
		[]int{1, 2, 3, 4, 5},
		desc,
		[]int{5, 4, 3, 2, 1},
	},
}

func TestSort(t *testing.T) {
	for i, e := range SortTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			clone := slices.Clone(e.input)
			slices.SortFunc(clone, e.fn)

			if !slices.Equal(clone, e.expected) {
				t.Fatalf("got %v, want %v", clone, e.expected)
			}
		})
	}
}
