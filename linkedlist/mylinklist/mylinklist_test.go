package mylinkedlist

import (
	"fmt"
	"slices"
	"strconv"
	"testing"
)

func TestConstructor(t *testing.T) {

	ml := Constructor()
	if ml.left.Next != ml.right || ml.right.Prev != ml.left {
		t.Fatalf("Constructor() = %v", ml)
	}

}

var AddAtTailTests = []struct {
	items []int
}{
	{[]int{1, 2, 3}},
	{Range(0, 100, 1)},
}

func TestAddAtTail(t *testing.T) {
	for i, e := range AddAtTailTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ll := Constructor()
			mll := &ll
			for i := range e.items {
				mll.AddAtTail(e.items[i])
			}
			asArray := gather(mll)
			if !slices.Equal(asArray, e.items) {
				t.Fatalf("%v != %v", asArray, e.items)
			}
		})
	}
}

func TestAddAtHead(t *testing.T) {
	for i, e := range AddAtTailTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ll := Constructor()
			mll := &ll
			for i := range e.items {
				mll.AddAtHead(e.items[i])
			}
			asArray := gather(mll)
			slices.Reverse(e.items)
			if !slices.Equal(asArray, e.items) {
				t.Fatalf("%v != %v", asArray, e.items)
			}
		})
	}
}

var InsertAtIndexTests = []struct {
	initial    []int
	index, val int
	expected   []int
}{
	{[]int{}, 0, 0, []int{0}},
	{[]int{}, 1, 0, []int{}},
	{[]int{1}, 0, 0, []int{0, 1}},
	{[]int{1, 2, 3}, 3, 4, []int{1, 2, 3, 4}},
	{[]int{1, 2, 3}, 1, 4, []int{1, 4, 2, 3}},
	{[]int{1, 2, 3}, 2, 4, []int{1, 2, 4, 3}},
}

func TestInsertAtIndex(t *testing.T) {
	for i, e := range InsertAtIndexTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			mll := FromArray(e.initial)

			mll.AddAtIndex(e.index, e.val)
			fmt.Println(mll)
			asArray := gather(mll)
			if !slices.Equal(asArray, e.expected) {
				t.Fatalf("%v != %v", asArray, e.expected)
			}
			if mll.Len() != len(e.expected) {
				t.Fatalf("mll.Len() = %v, want %v", mll.Len(), len(e.expected))
			}
		})
	}
}

var deleteAtIndexTests = []struct {
	initial  []int
	index    int
	expected []int
}{
	{},
	{[]int{0}, 0, []int{}},
	{[]int{0}, -1, []int{0}},
	{[]int{0}, 1, []int{0}},
	{[]int{0, 1, 2}, 0, []int{1, 2}},
	{[]int{0, 1, 2}, 2, []int{0, 1}},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, []int{0, 1, 2, 3, 4, 6, 7, 8, 9}},
}

func TestDeleteAtIndex(t *testing.T) {
	for i, e := range deleteAtIndexTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			mll := FromArray(e.initial)

			mll.DeleteAtIndex(e.index)
			fmt.Println(mll)
			asArray := gather(mll)
			if !slices.Equal(asArray, e.expected) {
				t.Fatalf("%v != %v", asArray, e.expected)
			}
		})
	}
}

func FromArray(arr []int) *MyLinkedList {
	ll := Constructor()
	mll := &ll
	for i := range arr {
		mll.AddAtTail(arr[i])
	}
	return mll
}

func gather(list *MyLinkedList) []int {
	acc := make([]int, 0)
	list.Visit(func(node *ListNode) bool {
		acc = append(acc, node.Val)
		return false
	})
	return acc
}

func Range(start, end, inc int) []int {
	if inc < 1 || start > end {
		return []int{}
	}
	rang := make([]int, 0, end-start)

	for i := 0; i < end; i++ {
		rang = append(rang, start+inc*i)
	}

	return rang
}

func Print(list *MyLinkedList) {
	list.Visit(func(node *ListNode) bool {
		fmt.Printf("%v", node)
		return false
	})
	fmt.Println()

}
