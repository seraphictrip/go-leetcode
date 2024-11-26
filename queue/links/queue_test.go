package links_test

import (
	"dsa/queue/links"
	"fmt"
	"slices"
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
			q := links.NewQueue[int]()

			for i := range e.inputs {
				q.Enqueue(e.inputs[i])
			}
			links.Print(q)
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

/*
The school cafeteria offers circular and square sandwiches at lunch break, referred to by numbers 0 and 1 respectively. All students stand in a queue. Each student either prefers square or circular sandwiches.

The number of sandwiches in the cafeteria is equal to the number of students. The sandwiches are placed in a stack. At each step:

If the student at the front of the queue prefers the sandwich on the top of the stack, they will take it and leave the queue.
Otherwise, they will leave it and go to the queue's end.
This continues until none of the queue students want to take the top sandwich and are thus unable to eat.

You are given two integer arrays students and sandwiches where sandwiches[i] is the type of the i​​​​​​th sandwich in the stack (i = 0 is the top of the stack) and students[j] is the preference of the j​​​​​​th student in the initial queue (j = 0 is the front of the queue). Return the number of students that are unable to eat.



Example 1:

Input: students = [1,1,0,0], sandwiches = [0,1,0,1]
Output: 0
Explanation:
- Front student leaves the top sandwich and returns to the end of the line making students = [1,0,0,1].
- Front student leaves the top sandwich and returns to the end of the line making students = [0,0,1,1].
- Front student takes the top sandwich and leaves the line making students = [0,1,1] and sandwiches = [1,0,1].
- Front student leaves the top sandwich and returns to the end of the line making students = [1,1,0].
- Front student takes the top sandwich and leaves the line making students = [1,0] and sandwiches = [0,1].
- Front student leaves the top sandwich and returns to the end of the line making students = [0,1].
- Front student takes the top sandwich and leaves the line making students = [1] and sandwiches = [1].
- Front student takes the top sandwich and leaves the line making students = [] and sandwiches = [].
Hence all students are able to eat.
Example 2:

Input: students = [1,1,1,0,0,1], sandwiches = [1,0,0,0,1,1]
11001 00011
10011 00011
00111 00011
..
..
111 011
Output: 3


Constraints:

1 <= students.length, sandwiches.length <= 100
students.length == sandwiches.length
sandwiches[i] is 0 or 1.
students[i] is 0 or 1.
*/

func countStudents(students []int, sandwiches []int) int {
	count := 0
	slices.Reverse(sandwiches)
	for len(students) >= 0 {
		student := students[0]                    // head of queue
		students = students[1:]                   // dequeue
		sandwich := sandwiches[len(sandwiches)-1] // top of stack

		if student == sandwich {
			sandwiches = sandwiches[0 : len(sandwiches)-1] // pop
			count = 0
		} else {
			students = append(students, student)
			count++
		}
		fmt.Printf("%v, %v\n", students, sandwiches)
		if count == len(students) {

			return count
		}
	}

	return 0
}

func countStudents2(students []int, sandwiches []int) int {
	pref := [2]int{0, 0}
	for i := range students {
		if students[i] == 0 {
			pref[0]++
		} else {
			pref[1]++
		}
	}

	for _, sand := range sandwiches {
		if pref[sand] > 0 {
			pref[sand]--
		} else {
			// return count of other
			sand ^= 1
			return pref[sand]
		}
		fmt.Println(pref)
	}
	return 0
}

func countStudents3(students []int, sandwhiches []int) int {
	res := len(students)
	// peference for 0 or 1 sandwhich
	preferences := [2]int{0, 0}

	for i := range students {
		preferences[students[i]]++
	}
	for i := range sandwhiches {
		if preferences[sandwhiches[i]] > 0 {
			preferences[sandwhiches[i]]--
			res--
		} else {
			break
		}

	}

	return res
}

var countStudentsTests = []struct {
	students   []int
	sandwiches []int
	expected   int
}{
	// {[]int{1, 1, 0, 0}, []int{0, 1, 0, 1}, 0},
	{[]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}, 3},
}

func TestCountStudents(t *testing.T) {
	for i, e := range countStudentsTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := countStudents3(e.students, e.sandwiches)
			if actual != e.expected {
				t.Fatalf("%v : %v", actual, e.expected)
			}
		})
	}
}

type Bits uint8

const (
	Flag0 Bits = 1 << iota
	Flag1
	Flag2
	Flag3
	Flag4
	Flag5
	Flag6
	Flag7
)

func Set(bs, flag Bits) Bits {
	return bs | flag
}

func Clear(bs, flag Bits) Bits {
	return bs &^ flag
}

func Toggle(bs, flag Bits) Bits {
	return bs ^ flag
}

func Has(bs, flag Bits) bool {
	return bs&flag != 0
}

/*
Implement a last-in-first-out (LIFO) stack using only two queues. The implemented stack should support all the functions of a normal stack (push, top, pop, and empty).

Implement the MyStack class:

void push(int x) Pushes element x to the top of the stack.
int pop() Removes the element on the top of the stack and returns it.
int top() Returns the element on the top of the stack.
boolean empty() Returns true if the stack is empty, false otherwise.
Notes:

You must use only standard operations of a queue, which means that only push to back, peek/pop from front, size and is empty operations are valid.
Depending on your language, the queue may not be supported natively. You may simulate a queue using a list or deque (double-ended queue) as long as you use only a queue's standard operations.


Example 1:

Input
["MyStack", "push", "push", "top", "pop", "empty"]
[[], [1], [2], [], [], []]
Output
[null, null, null, 2, 2, false]

Explanation
MyStack myStack = new MyStack();
myStack.push(1);
myStack.push(2);
myStack.top(); // return 2
myStack.pop(); // return 2
myStack.empty(); // return False


Constraints:

1 <= x <= 9
At most 100 calls will be made to push, pop, top, and empty.
All the calls to pop and top are valid.


Follow-up: Can you implement the stack using only one queue?

*/

type MyStack struct {
	q []int
}

func Constructor() MyStack {
	return MyStack{
		q: []int{},
	}
}

// Both stack and queue append to end, so this seems right...
func (s *MyStack) Push(x int) {
	// push and endqueue are same, add to end
	s.q = append(s.q, x)
}

func (s *MyStack) Pop() int {
	n := len(s.q)
	item := -1
	for i := 0; i < n; i++ {
		item = s.q[0] // peek
		s.q = s.q[1:] // dequeue
		if i < n-1 {
			s.q = append(s.q, item) // enqueue
		}
	}
	return item
}

func (s *MyStack) Top() int {
	return s.q[len(s.q)-1]
}

func (s *MyStack) Empty() bool {
	return len(s.q) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func TestMyStack(t *testing.T) {
	stack := Constructor()
	if !stack.Empty() {
		t.Fatal("Empty()")
	}
	stack.Push(0)
	if stack.Empty() {
		t.Fatal("Empty() after push")
	}
	if stack.Top() != 0 {
		t.Fatalf("Top() failed")
	}
	item := stack.Pop()
	if item != 0 {
		t.Fatal("Pop() wrong")
	}
	if !stack.Empty() {
		t.Fatal("should be Empty() after pop all")
	}

	pushes := []int{0, 1, 2}
	for i := range pushes {
		stack.Push(pushes[i])
	}
	slices.Reverse(pushes)
	pops := pushes
	for i := range pops {
		if stack.Pop() != pops[i] {
			t.Fatalf("Pop() failed at %v", pops[i])
		}
	}

}
