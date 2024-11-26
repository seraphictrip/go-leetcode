package matrix_test

import (
	"container/list"
	"dsa/matrix"
	"fmt"
	"strconv"
	"testing"
)

var matrixTests = []struct {
	grid     [][]int
	expected int
}{
	{[][]int{{0, 0, 0, 0}}, 1},
	{[][]int{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}, 2},
}

func TestMatrix(t *testing.T) {
	for i, e := range matrixTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := matrix.CountPaths(e.grid)
			if actual != e.expected {
				t.Fatalf("CountPaths(%v) = %v, want %v", e.grid, actual, e.expected)
			}
		})
	}
}

func TestQueue(t *testing.T) {
	queue := list.New()

	queue.PushBack([]int{0, 0})
	queue.PushBack([]int{0, 1})
	front := queue.Front()
	fmt.Println(front.Value, queue.Len())
	queue.Remove(front)
	fmt.Println(front.Value, queue.Len())

}
