package algo

func QuickSort(arr []int) {
	// invariant everything

	var quicksort func(nums []int, l, r int)

	quicksort = func(nums []int, l, r int) {
		part := partition(nums, l, r)
		if part == -1 {
			return
		}
		quicksort(nums, l, part-1)
		quicksort(nums, part+1, r)
	}

	quicksort(arr, 0, len(arr)-1)

}

func partition(nums []int, l, r int) int {
	if l >= r {
		return -1
	}
	pivot := nums[r]
	part := l
	for i := l; i < r; i++ {
		if nums[i] <= pivot {
			swap(nums, part, i)
			part++
		}
	}
	// swap pivot into place
	swap(nums, part, r)
	return part
}
