package sort

import "fmt"

// for each position find correct item
// invariant, every thing left of i is sorted

// O(n^2)
// inplace
func SimpleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[i] {
				swap(arr, i, j)
			}
		}
	}
	return arr
}

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// Insertion sort
func InsertionSort(arr []int) []int {
	// partition so everything to left of i is in sorted order
	//[5...]
	// [1,5...]
	//[1,3,5...]
	// [1,3,4,5..]
	for i := 1; i < len(arr); i++ {
		j := i - 1
		// insert into place by swapping down until in right position
		for j >= 0 && arr[j] > arr[j+1] {
			swap(arr, j+1, j)
			j--
		}

	}
	return arr
}

func InsertionSort2(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		for j >= 0 && arr[j+1] < arr[j] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
			j--
		}
	}
	return arr
}

func InsertionSort3(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		for j >= 0 && arr[j+1] < arr[j] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
			j--
		}
	}
	return arr
}

func InsertionSortRecursive(arr []int) []int {
	return insert(arr, 1)
}

func insert(arr []int, i int) []int {
	if i >= len(arr) {
		return arr
	}
	j := i - 1
	for j >= 0 && arr[j] > arr[j+1] {
		swap(arr, j, j+1)
		j--
	}
	return insert(arr, i+1)
}

func MergeSort(arr []int) []int {
	return mergeSort(arr, 0, len(arr)-1)
}

func merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}

	}

	for i < len(a) {
		result = append(result, a[i])
		i++
	}
	for j < len(b) {
		result = append(result, b[j])
		j++
	}

	return result
}

func mergeSort(arr []int, start, end int) []int {
	if end-start+1 <= 1 {
		return arr
	}
	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)

	return merge2(arr, start, mid, end)
}

func merge2(arr []int, start, mid, end int) []int {
	i := start
	j := mid + 1
	curr := start
	a := make([]int, mid-start+1)
	b := make([]int, end-mid+1)
	copy(a, arr[start:mid+1])
	copy(b, arr[mid+1:])

	for i <= mid && j <= end {
		if arr[i] < arr[j] {
			arr[curr] = a[i]
			i++
		} else {
			arr[curr] = b[j]
			j++
		}
		curr++
	}

	for i <= mid {
		swap(arr, i, curr)
		i++
		curr++
	}

	for j <= end {
		swap(arr, j, curr)
		j++
		curr++
	}

	return arr
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	if n == nil {
		return "nil"
	}
	str := fmt.Sprintf("{%v}->%v", n.Val, n.Next)
	return str
}

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	head := lists[0]
	for i := 1; i < len(lists); i++ {
		head = mergeLL(head, lists[i])
	}
	return head
}

func mergeLL(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	var head *ListNode = nil
	var cur *ListNode = nil
	if a.Val < b.Val {
		head = a
		a = a.Next
	} else {
		head = b
		b = b.Next
	}
	cur = head
	for a != nil && b != nil {
		if a.Val < b.Val {
			cur.Next = a
			a = a.Next
		} else {
			cur.Next = b
			b = b.Next
		}
		cur = cur.Next
	}

	if a != nil {
		cur.Next = a
	}
	if b != nil {
		cur.Next = b
	}

	return head
}

func ToLinkedList(arr []int) *ListNode {
	n := len(arr)

	if n == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	cur := head
	for i := 1; i < n; i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
	return head
}

func FromLinkedList(acc []int, node *ListNode) []int {
	if node == nil {
		return acc
	}
	acc = append(acc, node.Val)
	return FromLinkedList(acc, node.Next)
}

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, start, end int) {
	if end-start <= 0 {
		return
	}
	partitionIndex := partition(arr, start, end)
	quickSort(arr, start, partitionIndex-1)
	quickSort(arr, partitionIndex+1, end)
}

func partition(arr []int, start, end int) int {
	pivotValue := arr[end]
	partitionIndex := start

	for i := start; i < end; i++ {
		if arr[i] < pivotValue {
			swap(arr, partitionIndex, i)
			partitionIndex++
		}
	}

	swap(arr, partitionIndex, end)

	return partitionIndex
}
