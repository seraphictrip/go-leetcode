package stack_test

import (
	"dsa/stack"
	"fmt"
	"strconv"
	"testing"
)

var StackTests = []struct {
	inputs []int
}{
	{},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8}},
	{[]int{5, 4, 3, 2, 1}},
}

func TestStack(t *testing.T) {
	for i, e := range StackTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := stack.New[int]()
			for i, input := range e.inputs {
				s.Push(input)
				top, err := s.Top()
				if err != nil {
					t.Fatalf("Top() = %v, %v, want %v, %v", top, err, input, nil)
				}
				if s.Len() != i+1 {
					t.Fatalf("s.Len() = %v, want %v", s.Len(), i+1)
				}
				fmt.Println(gather(s), s)
			}
		})
	}
}

func TestMonoIncStack(t *testing.T) {
	for i, e := range StackTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := stack.NewMonotonicIncreasingStack[int]()
			for _, input := range e.inputs {
				s.Push(input)
				top, err := s.Top()
				if err != nil {
					t.Fatalf("Top() = %v, %v, want %v, %v", top, err, input, nil)
				}
				// if s.Len() != i+1 {
				// 	t.Fatalf("(%v).Len() = %v, want %v", s, s.Len(), i+1)
				// }
				fmt.Println(gather(s.Stack), s)
			}
		})
	}
}

func TestMonoDecStack(t *testing.T) {
	for i, e := range StackTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := stack.NewMonotonicDecreasingStack[int]()
			for _, input := range e.inputs {
				s.Push(input)
				top, err := s.Top()
				if err != nil {
					t.Fatalf("Top() = %v, %v, want %v, %v", top, err, input, nil)
				}
				// if s.Len() != i+1 {
				// 	t.Fatalf("(%v).Len() = %v, want %v", s, s.Len(), i+1)
				// }
				fmt.Println(gather(s.Stack), s)
			}
		})
	}
}

// Time Complexity O(n)
// Space Complexity O(1), input/output value space doesn't count
func gather[T any](s *stack.Stack[T]) []T {
	result := make([]T, s.Len())
	// take each item off top and put in array
	for i := 0; !s.IsEmpty(); i++ {
		result[i], _ = s.Top()
		s.Pop()
	}
	// put items back on stack in reverse order
	for i := len(result) - 1; i >= 0; i-- {
		s.Push(result[i])
	}
	return result
}

/*
Create (hypothesis, novel knowledge generation)
Evaluate (Judgment "W" questions, unlocks Prioritize )
Analyze (contextual/compare and contrast, venn diagrams, tables, summaries, mindmaps)
* Compare constrast
taxonomy level <Level>
Apply (solve 1 to 1 problems)
Understand (Explain)
Remember (re-re remember/repeat/regurgitate)



PROMPT: give me questions at <education level> for <subject> at Blooms Revised
*/

// NextGreaterElem returns the index next greater element to the right
// of i, or -1

func NextGreaterElem(arr []int, i int) int {
	if i > len(arr)-2 {
		// can't overrun array, or find greater past last elem
		return -1
	}
	curr := arr[i]
	for i := i + 1; i < len(arr); i++ {
		if arr[i] > curr {
			return i
		}
	}
	return -1
}

/*
Input: arr[] = [ 4 , 5 , 2 , 25 ]
Output:  4      –>   5 (1)

	 5      –>   25 (3)
	 2      –>   25 (3)
	25     –>   -1
*/
var NGETests = []struct {
	arr      []int
	i        int
	expected int
}{
	{nil, 0, -1},
	{[]int{4, 5, 2, 25}, 0, 1},
	{[]int{4, 5, 2, 25}, 1, 3},
	{[]int{4, 5, 2, 25}, 2, 3},
	{[]int{4, 5, 2, 25}, 3, -1},
}

func TestNGE(t *testing.T) {
	for i, e := range NGETests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := NextGreaterElem(e.arr, e.i)

			if actual != e.expected {
				t.Fatalf("NextGreaterElem(%v, %v) = %v, want %v", e.arr, e.i, actual, e.expected)
			}
		})
	}
}
