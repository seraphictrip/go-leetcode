package linkedlist_test

import (
	"dsa/linkedlist"
	"fmt"
	"math"
	"math/rand"
	"slices"
	"strconv"
	"testing"
)

var reverseTests = []struct {
	head *linkedlist.ListNode
}{
	// {},
	// {
	// 	head: &linkedlist.ListNode{Val: 4},
	// },
	{
		head: fromArray([]int{1, 2, 3, 4, 5}),
	},
}

func TestReverse(t *testing.T) {
	for i, e := range reverseTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			head := linkedlist.ReverseList2(e.head)
			fmt.Println(head)
		})
	}
}

func fromArray(vals []int) *linkedlist.ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &linkedlist.ListNode{
		Val: vals[0],
	}
	head.Next = fromArray(vals[1:])
	return head
}

func Range(start, end, inc int) []int {
	count := end - start
	result := make([]int, count)
	for i := 0; i < count; i++ {
		result[i] = start + (inc * i)
	}
	return result
}

// just slices.Reverese but with return value for chaining
func Reverse(s []int) []int {
	slices.Reverse(s)
	return s
}

func TestLinkedListBasics(t *testing.T) {
	for i := 0; i <= 6; i++ {
		from := Range(0, i, 1)
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ll := linkedlist.FromArray(from)
			fmt.Println(ll)
			list := ll.ToArray()
			if !slices.Equal(list, from) {
				t.Fatalf("ll.ToArray() = %v, want, %v", list, from)
			}
		})
	}
}

var AppendTests = []struct {
	initial  *linkedlist.LinkedList
	appends  []int
	expected *linkedlist.LinkedList
}{
	{},
	{linkedlist.FromArray([]int{1, 2, 3}), nil, linkedlist.FromArray([]int{1, 2, 3})},
	{new(linkedlist.LinkedList), []int{1, 2, 3}, linkedlist.FromArray([]int{1, 2, 3})},
	{new(linkedlist.LinkedList), Range(0, 100, 1), linkedlist.FromArray(Range(0, 100, 1))},
	{linkedlist.FromArray([]int{1, 2, 3}), Range(0, 100, 1), linkedlist.FromArray(append([]int{1, 2, 3}, Range(0, 100, 1)...))},
}

func TestAppend(t *testing.T) {
	for i, e := range AppendTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ll := e.initial
			for i := range e.appends {
				ll.Append(e.appends[i])
			}

			if !ll.Equals(e.expected) {
				t.Fatalf("%v != %v", ll, e.expected)
			}
		})
	}
}

var PrependTests = []struct {
	initial  *linkedlist.LinkedList
	prepends []int
	expected *linkedlist.LinkedList
}{
	{},
	{linkedlist.FromArray([]int{1, 2, 3}), nil, linkedlist.FromArray([]int{1, 2, 3})},
	{new(linkedlist.LinkedList), []int{1, 2, 3}, linkedlist.FromArray([]int{3, 2, 1})},
	{new(linkedlist.LinkedList), Range(0, 100, 1), linkedlist.FromArray(Reverse(Range(0, 100, 1)))},
	{linkedlist.FromArray([]int{1, 2, 3}), Range(0, 100, 1), linkedlist.FromArray(append(Reverse(Range(0, 100, 1)), []int{1, 2, 3}...))},
}

func TestPrepend(t *testing.T) {
	for i, e := range PrependTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ll := e.initial
			for i := range e.prepends {
				ll.Prepend(e.prepends[i])
			}

			if !ll.Equals(e.expected) {
				t.Fatalf("%v != %v", ll, e.expected)
			}
		})
	}
}

func haystack(size, needle int) []int {
	hs := Range(0, size, 1)
	needleIndex := rand.Intn(size)
	hs[needleIndex] = needle
	return hs
}

var DeleteTests = []struct {
	initial  *linkedlist.LinkedList
	val      int
	expected *linkedlist.LinkedList
}{
	{new(linkedlist.LinkedList), 100, new(linkedlist.LinkedList)},
	{linkedlist.Just(1), 1, linkedlist.Empty()},
	{linkedlist.FromArray([]int{1, 2, 3, 4, 5}), 1, linkedlist.FromArray([]int{2, 3, 4, 5})},
	{linkedlist.FromArray([]int{1, 2, 3, 4, 5}), 5, linkedlist.FromArray([]int{1, 2, 3, 4})},
	{linkedlist.FromArray([]int{1, 2, 3, 4, 5}), 100, linkedlist.FromArray([]int{1, 2, 3, 4, 5})},
	{linkedlist.FromArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), 5, linkedlist.FromArray([]int{1, 2, 3, 4, 6, 7, 8, 9})},
}

func TestDelete(t *testing.T) {
	for i, e := range DeleteTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ll := e.initial
			ll.Delete(e.val)

			if !ll.Equals(e.expected) {
				t.Fatalf("%v != %v", ll, e.expected)
			}
		})
	}
}

var DeletePercisionTests = []struct {
	haystack *linkedlist.LinkedList
	needle   int
	expected bool
}{
	{
		new(linkedlist.LinkedList), 0, false,
	},
	{linkedlist.FromArray(haystack(1000, -99)), -99, true},
	{linkedlist.FromArray(haystack(10000, -99)), -99, true},
	{linkedlist.FromArray(haystack(1000000, -99)), -99, true},
	{linkedlist.FromArray(haystack(math.MaxInt16, -99)), -99, true},
}

func TestDeletePercision(t *testing.T) {
	for i, e := range DeletePercisionTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := e.haystack.Delete(e.needle)
			if actual != e.expected {
				t.Fatalf("ll.Delete(%v) = %v, want %v", e.needle, actual, e.expected)
			}
		})
	}
}

func TestMyLinkedList(t *testing.T) {
	//  for i, e := range MyLinkedListTests {
	// 	e := e
	// 	t.Run(strconv.Itoa(i), func(t *testing.T){

	// 	})
	// }
	mll := linkedlist.Constructor()
	fmt.Println(mll)
}
