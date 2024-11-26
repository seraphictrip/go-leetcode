package algo

import (
	"slices"
)

func MergeSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	mid := n / 2
	return merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
}

func merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
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

func MergeSortInplace(arr []int) []int {
	var mergesort func(arr []int, start, end int) []int

	mergesort = func(arr []int, start, end int) []int {
		if end-start+1 <= 1 {
			return arr
		}
		mid := (start + end) / 2
		mergesort(arr, start, mid)
		mergesort(arr, mid+1, end)
		mergeinplace(arr, start, mid, end)
		return arr
	}
	return mergesort(arr, 0, len(arr)-1)
}

func mergeinplace(arr []int, start, mid, end int) {
	L := slices.Clone(arr[start : mid+1])
	R := slices.Clone(arr[mid+1 : end+1])

	i := 0
	j := 0
	cur := start
	for i < len(L) && j < len(R) {
		if L[i] < R[j] {
			arr[cur] = L[i]
			i++
		} else {
			arr[cur] = R[j]
			j++
		}
		cur++
	}

	for i < len(L) {
		arr[cur] = L[i]
		i++
		cur++
	}
	for j < len(R) {
		arr[cur] = R[j]
		j++
		cur++
	}
}
