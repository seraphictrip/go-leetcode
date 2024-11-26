package algo_test

import (
	"dsa/sort/algo"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"testing"
)

var SortTests = []struct {
	arr []int
}{
	{[]int{2, 4, 1, 3, 7, 6, 5}},
	{[]int{2, 4, 1, 3, 3}},
}

func TestInsertionSort(t *testing.T) {
	for i, e := range SortTests {
		e := e
		original := slices.Clone(e.arr)
		expected := slices.Clone(e.arr)
		slices.Sort(expected)
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			algo.InsertionSort(e.arr)
			if !slices.Equal(e.arr, expected) {
				t.Fatalf("InsertionSort(%v) = %v, want %v", original, e.arr, expected)
			}

		})
	}
}

func TestMergeSort(t *testing.T) {
	for i, e := range SortTests {
		e := e
		original := slices.Clone(e.arr)
		expected := slices.Clone(e.arr)
		slices.Sort(expected)
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := algo.MergeSortInplace(e.arr)
			if !slices.Equal(actual, expected) {
				t.Fatalf("MergeSort(%v) = %v, want %v", original, actual, expected)
			}

		})
	}
}

func TestQuickSort(t *testing.T) {
	for i, e := range SortTests {
		e := e
		original := slices.Clone(e.arr)
		expected := slices.Clone(e.arr)
		slices.Sort(expected)
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			algo.QuickSort(e.arr)
			if !slices.Equal(e.arr, expected) {
				t.Fatalf("MergeSort(%v) = %v, want %v", original, e.arr, expected)
			}

		})
	}
}

func parseTime(s string) []string {
	// 07:05:45PM
	regex := regexp.MustCompile(`(\d{2}):(\d{2}):(\d{2})(.*)`)
	return regex.FindStringSubmatch(s)
}

var funcTests = []struct {
	s string
}{
	{"07:05:45PM"},
}

func TestFunc(t *testing.T) {
	for i, e := range funcTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			fmt.Println(parseTime(e.s))
		})
	}
}
