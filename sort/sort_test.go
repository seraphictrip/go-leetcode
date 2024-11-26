package sort

import (
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"testing"
)

var sortTests = []struct {
	input []int
}{
	{[]int{5, 1, 3, 4, 2}},
	{[]int{5, 4, 3, 2, 1}},
	{shuffle(Range(0, 1000, 1))},
}

func TestSimpleSort(t *testing.T) {
	for i, e := range sortTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			clone := slices.Clone(e.input)
			actual := SimpleSort(e.input)
			slices.Sort(clone)
			if !slices.Equal(clone, actual) {
				t.Fatalf("InsertionSort(...)= %v, want %v", actual, clone)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for i, e := range sortTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			clone := slices.Clone(e.input)
			actual := InsertionSort(e.input)
			slices.Sort(clone)
			if !slices.Equal(clone, actual) {
				t.Fatalf("InsertionSort(...)= %v, want %v", actual, clone)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for i, e := range sortTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			clone := slices.Clone(e.input)
			actual := MergeSort(e.input)
			slices.Sort(clone)
			if !slices.Equal(clone, actual) {
				t.Fatalf("MergeSort(...)= %v, want %v", actual, clone)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	for i, e := range sortTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			clone := slices.Clone(e.input)
			QuickSort(e.input)
			slices.Sort(clone)
			if !slices.Equal(clone, e.input) {
				t.Fatalf("QuickSort(...)= %v, want %v", e.input, clone)
			}
		})
	}
}

func TestBucketSort(t *testing.T) {
	for i, e := range sortTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			clone := slices.Clone(e.input)
			codomain := make([]int, len(e.input)+1)
			for i := range codomain {
				codomain[i] = i
			}
			BucketSortGeneric(e.input, func(x int) int { return x }, func(x int) int { return x }, codomain)
			slices.Sort(clone)
			if !slices.Equal(clone, e.input) {
				t.Fatalf("BucketSort(...)= %v, want %v", e.input, clone)
			}
		})
	}
}

func Range(start, count, inc int) []int {
	arr := make([]int, count)
	for i := 0; i < count; i++ {
		arr[i] = start + (i * inc)
	}
	return arr
}

func shuffle(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		j := rand.Intn(len(arr))
		swap(arr, i, j)
	}
	return arr
}

var mergeKTests = []struct {
	lists []*ListNode
}{
	{},
	{[]*ListNode{ToLinkedList([]int{1, 2, 4}), ToLinkedList([]int{3, 5, 6})}},
	{[]*ListNode{
		ToLinkedList([]int{1, 2, 4}),
		ToLinkedList([]int{3, 5, 6}),
		ToLinkedList([]int{0, 8, 10}),
	}},
}

func TestMergeK(t *testing.T) {
	for i, e := range mergeKTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			fmt.Println(MergeKLists(e.lists))
		})
	}
}

var ToLinkedListTests = []struct {
	arr []int
}{
	{},
	{[]int{0}},
	{[]int{0, 1}},
	{[]int{0, 1, 2, 3, 4, 5}},
}

func TestToLinkedList(t *testing.T) {
	for i, e := range ToLinkedListTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ll := ToLinkedList(e.arr)
			arr := FromLinkedList(nil, ll)
			if !slices.Equal(e.arr, arr) {
				t.Fatalf("want %v, got %v", e.arr, arr)
			}
		})
	}
}
