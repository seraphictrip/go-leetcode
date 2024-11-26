package functional_test

import (
	"dsa/patterns/functional"
	"slices"
	"strconv"
	"testing"
)

var PartitionTests = []struct {
	inputs   []int
	choice   func(int) bool
	expected []int
}{
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		func(n int) bool { return n%2 == 0 },
		[]int{2, 4, 6, 8, 1, 3, 5, 7, 9},
	},
	{
		Range(100),
		func(n int) bool { return n%2 == 0 },
		[]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 41, 43, 45, 47, 49, 51, 53, 55, 57, 59, 61, 63, 65, 67, 69, 71, 73, 75, 77, 79, 81, 83, 85, 87, 89, 91, 93, 95, 97, 99},
	},
	{
		Range(1000),
		func(n int) bool { return n%2 == 0 },
		[]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 41, 43, 45, 47, 49, 51, 53, 55, 57, 59, 61, 63, 65, 67, 69, 71, 73, 75, 77, 79, 81, 83, 85, 87, 89, 91, 93, 95, 97, 99},
	},
}

func Range(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i
	}
	return result
}

func TestPartition(t *testing.T) {
	for i, e := range PartitionTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			functional.PartitionStable(e.inputs, e.choice)
			// TODO: change this, stable is not garunteed in this version
			// will just use ParitionStable for now
			if !slices.Equal(e.inputs, e.expected) {
				t.Fatalf("Paritition(...) = %v, want %v", e.inputs, e.expected)
			}
		})
	}
}

var FilterTests = []struct {
	inputs   []int
	choice   func(int) bool
	expected []int
}{
	{},
	{[]int{1}, func(n int) bool { return n == 1 }, []int{1}},
	{[]int{1}, func(n int) bool { return n != 1 }, []int{}},
	{Range(100), func(n int) bool { return n%2 == 0 }, []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98}},
}

func TestFilter(t *testing.T) {
	for i, e := range FilterTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			e := e
			t.Run(strconv.Itoa(i), func(t *testing.T) {
				result := functional.Filter(e.inputs, e.choice)
				if !slices.Equal(result, e.expected) {
					t.Fatalf("Paritition(%v) = %v, want %v", e.inputs, result, e.expected)
				}
			})
		})
	}
}

var MapTests = []struct {
	inputs    []int
	transform func(n int) int
	expected  []int
}{
	{
		inputs:    []int{},
		transform: func(n int) int { return n },
		expected:  []int{},
	},
	{
		inputs:    []int{1, 2, 3, 4, 5},
		transform: functional.Identity[int],
		expected:  []int{1, 2, 3, 4, 5},
	},
	{
		inputs:    Range(100),
		transform: functional.Identity[int],
		expected:  Range(100),
	},
	{
		inputs:    []int{1, 2, 3, 4, 5},
		transform: func(n int) int { return n * n },
		expected:  []int{1, 4, 9, 16, 25},
	},
	{
		inputs:    Range(100),
		transform: func(n int) int { return 0 },
		expected:  make([]int, 100),
	},
}

func TestMap(t *testing.T) {
	for i, e := range MapTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := functional.Map(e.inputs, e.transform)
			if !slices.Equal(result, e.expected) {
				t.Fatalf("Paritition(%v) = %v, want %v", e.inputs, result, e.expected)
			}
		})
	}
}

var ReduceTests = []struct {
	inputs   []int
	reducer  func(acc, cur int) int
	initial  int
	expected int
}{
	{nil, func(acc, n int) int { return acc + n }, 0, 0},
	{nil, func(acc, n int) int { return acc + n }, 100, 100},
	{[]int{1, 2, 3, 4, 5}, func(acc, n int) int { return acc + n }, 0, 15},
	{[]int{1, 2, 3, 4, 5}, func(acc, n int) int { return acc * n }, 1, (1 * 1 * 2 * 3 * 4 * 5)},
	{[]int{1, 2, 3, 4, 5}, func(acc, n int) int { return acc - n }, 100, (100 - 1 - 2 - 3 - 4 - 5)},
}

func TestReduce(t *testing.T) {
	for i, e := range ReduceTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := functional.Reduce(e.inputs, e.reducer, e.initial)
			if actual != e.expected {
				t.Fatalf("Reduce(%v, reducer, %v) = %v, want %v", e.inputs, e.initial, actual, e.expected)
			}
		})
	}
}
