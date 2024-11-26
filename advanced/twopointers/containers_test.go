package twopointers_test

import (
	"dsa/advanced/twopointers"
	"strconv"
	"testing"
)

var MaxAreaTests = []struct {
	arr      []int
	expected int
}{
	{},
	// [1,8,6,2,5,4,8,3,7]
	//
	//	L				R		area: 8 maxarea: 8 (8*1)
	//	 L				R		area: 49 maxarea: 49 (7*7)
	//	 L			  R	 		area: 18 maxarea: 49
	//	 L          R			area: 40 maxarea: 49
	//	 L        R				area: 16 maxarea: 49
	//	 L      R				area: 15 maxarea: 49
	{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
}

func TestMaxArea(t *testing.T) {
	for i, e := range MaxAreaTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := twopointers.MaxArea(e.arr)
			if actual != e.expected {
				t.Fatalf("MaxArea(%v) = %v, want %v", e.arr, actual, e.expected)
			}
		})
	}
}

var trappedTests = []struct {
	height   []int
	expected int
}{
	// {[]int{4, 2, 0, 3, 2, 5}, 9},
	//  L L L L				R  R
	// [0,1,0,2,1,0,1,3,2,1,2,1]
	{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
}

func TestTrapped(t *testing.T) {
	for i, e := range trappedTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := twopointers.Trap(e.height)
			if actual != e.expected {
				t.Fatalf("Trap(%v) = %v, want %v", e.height, actual, e.expected)
			}
		})
	}
}

// L  R
// [0,1,0,2,1,0,1,3,2,1,2,1]
// func Trap2(height []int) int {
// 	trapped := 0
// 	L, R := 0, len(height)-1
// 	maxL, maxR := height[L], height[R]

// 	for L < R {
// 		curL, curR := height[L], height[R]
// 		minwall := min(maxL, maxR)
// 		if curL < curR

// 	}
// 	return 0
// }
