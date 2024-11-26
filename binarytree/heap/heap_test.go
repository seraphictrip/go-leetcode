package heap_test

import (
	"dsa/binarytree/heap"
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"testing"
)

var PushTests = []struct {
	inputs []int
}{
	// {},
	// {[]int{1}},
	{[]int{10, 5, 6, 7, 1, 9}},
	{shuffle(Seq(0, 10, 1))},
}

func TestPush(t *testing.T) {
	for i, e := range PushTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			h := &heap.MinHeap[int]{}
			for _, num := range e.inputs {
				h.Push(num)
			}
			fmt.Println(h)
		})
	}
}

var ParentIndexTests = []struct {
	i, expected int
}{
	// (0-1)/2 = 0 // int math
	{0, 0},
	// (1-1)/2 = 0
	{1, 0},
	// (2-1)/2 = 0
	{2, 0},
	// (3-1)/2 = 1
	{3, 1},
	// (4-1)/2 = 1
	{4, 1},
	// (5-1)/2 = 2
	{5, 2},
	// (6-1)/2 = 2
	{6, 2},
	// (7-1)/2 = 3
	{7, 3},
	// (8-1)/2 = 3
	{8, 3},
	// (9-1)/2 = 4
	{9, 4},
	// (9-1)/2 = 4
	{10, 4},
}

func TestParentIndex(t *testing.T) {
	for i, e := range ParentIndexTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := heap.ParentIndex(e.i)
			if actual != e.expected {
				t.Fatalf("your base understanding is low")
			}
		})
	}
}

var LeftChildIndexTests = []struct {
	i, expected int
}{
	// 2i+1, though if used sentinel could be 2i
	// OBSERVATION: It is interesting the distance between parent and left child inc by 1
	// this is same outcome of just doubling when 1 based, I never really though about it
	// 2*0+1 = 1
	{0, 1}, // 1
	// 2*1+1 = 3
	{1, 3}, // 2
	{2, 5}, // 3
	{3, 7}, // 4
	{4, 9}, // 5
	{5, 11},
	{6, 13},
	{7, 15},
	{8, 17},
	{9, 19},
	{10, 21},
}

func TestLeftChildIndex(t *testing.T) {
	for i, e := range LeftChildIndexTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := heap.LeftChildIndex(e.i)
			if actual != e.expected {
				t.Fatalf("your base understanding is low")
			}
		})
	}
}

var RightChildIndexTests = []struct {
	i, expected int
}{
	// 2i+2, though if used sentinel could be 2i + 1
	// 2*0+2 = 2
	{0, 2},
	// 2*1+2 = 4
	{1, 4},
	// 2*2 + 2 = 6
	{2, 6},
	{3, 8},
	{4, 10},
	{5, 12},
	{6, 14},
	{7, 16},
	{8, 18},
	{9, 20},
	{10, 22},
}

func TestRightChildIndex(t *testing.T) {
	for i, e := range RightChildIndexTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := heap.RightChildIndex(e.i)
			if actual != e.expected {
				t.Fatalf("RightChildIndex(%v) = %v, want %v", e.i, actual, e.expected)
			}
		})
	}
}

// var PopTests = []struct {
// }{}

// func TestPop(t *testing.T) {
// 	for i, e := range PopTests {
// 		e := e
// 		t.Run(strconv.Itoa(i), func(t *testing.T) {
// 			// code
// 		})
// 	}
// }

var MinHeapifyNaiveTests = []struct {
	inputs   []int
	expected []int
}{
	{},
	// (1)
	{[]int{1}, []int{1}},
	// 			(1)
	//		(2)
	{[]int{1, 2}, []int{1, 2}},
	//	(2)    =>		(1)
	//				 (2)
	{[]int{2, 1}, []int{1, 2}},
	// 		(1)
	//	(2)		(3)
	{[]int{1, 2, 3}, []int{1, 2, 3}},
	//	(3) => 	(2)				(1)
	//		 (3)		=>	(3)		(2)
	{[]int{3, 2, 1}, []int{1, 3, 2}},
	//			(1)
	//		(2)		(3)
	//	(4)
	{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
	// 	(4) =>	(3)				(2)				(1)
	//		 (4)		=> 	 (4)	(3) =>	(2)		(3)
	//									  (4)
	{[]int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
	{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	//	(5) => 	(4)			(3)			(2)			(1)
	//		 (5)	=>   (5)	(4) => (3)	(4) => (2)	(4)
	//								  (5)		(5)  (3)
	{[]int{5, 4, 3, 2, 1}, []int{1, 2, 4, 5, 3}},
	//	(2) => (2)		(1)			(1)				(1)
	//		 (4)	=> (4) (2) => (4)	(2) => 	  (3)	(2)
	//							(5)				(5) (4)
	//
	{[]int{2, 4, 1, 5, 3}, []int{1, 3, 2, 5, 4}},
}

func TestMinHeapifyNaive(t *testing.T) {
	for i, e := range MinHeapifyNaiveTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			h := heap.MinHeapifyNaive[int](e.inputs)
			if !slices.Equal(h.ToArray(), e.expected) {
				t.Fatalf("%v != %v", h.ToArray(), e.expected)
			}
		})
	}
}

func Seq(start, count, inc int) []int {
	result := make([]int, 0, count)
	for i := 0; i < count; i++ {
		result = append(result, start+(i*inc))
	}
	return result
}

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func shuffle(arr []int) []int {
	for i := range arr {
		swap(arr, i, rand.Intn(len(arr)))
	}
	return arr
}
