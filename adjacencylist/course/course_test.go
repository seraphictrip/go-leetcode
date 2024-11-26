package course_test

import (
	"dsa/adjacencylist/course"
	"slices"
	"strconv"
	"testing"
)

var CanFinishTests = []struct {
	numCourses    int
	prerequisites [][]int
	expected      bool
}{
	{2, [][]int{{1, 0}}, true},
	{2, [][]int{{1, 0}, {0, 1}}, false},
	{20, [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}}, false},
}

func TestCanFinish(t *testing.T) {
	for i, e := range CanFinishTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := course.CanFinish(e.numCourses, e.prerequisites)
			if actual != e.expected {
				t.Fatalf("CanFinish(%v, %v) = %v, want %v", e.numCourses, e.prerequisites, actual, e.expected)
			}
		})
	}
}

var FindOrderTests = []struct {
	numCourses    int
	prerequisites [][]int
	expected      []int
}{
	// {2, [][]int{{1, 0}}, []int{0, 1}},
	{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, []int{0, 2, 1, 3}},
}

func TestFindOrder(t *testing.T) {
	for i, e := range FindOrderTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// code
			actual := course.FindOrder(e.numCourses, e.prerequisites)
			if !slices.Equal(actual, e.expected) {
				t.Fatalf("FindOrder(%v, %v) = %v, want %v", e.numCourses, e.prerequisites, actual, e.expected)
			}
		})
	}
}
