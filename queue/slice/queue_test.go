package slice_test

import (
	"dsa/queue/slice"
	"strconv"
	"testing"
)

var QueueTests = []struct {
	inputs []int
}{
	{[]int{1, 2, 3, 4, 5, 6}},
	{Range(0, 10, 1)},
	{Range(0, 100, 1)},
	{Range(0, 10000, 2)},
	{Range(0, 10000, 10)},
}

func Range(start, count, inc int) []int {
	result := make([]int, count)
	for i := range result {
		result[i] = start + (i * inc)
	}
	return result
}

func TestQueue(t *testing.T) {
	for i, e := range QueueTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			q := slice.NewQueue[int]()

			for i := range e.inputs {
				q.Enqueue(e.inputs[i])
			}

			for i := range e.inputs {
				actual := q.Dequeue()
				expected := e.inputs[i]

				if actual != expected {
					t.Fatalf("q.Dequeue() = %v, want %v", actual, expected)
				}
			}

			if !q.IsEmpty() {
				t.Fatalf("%v not empty", q)
			}
		})
	}
}
