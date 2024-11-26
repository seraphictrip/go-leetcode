package binarysearch

// [0,1,2,3,4,5,6] 1
// mid = (0+7)/2 = 3, inputs[mid]=3, target < inputs[mid]
// right = 3-1
func BinarySearch(inputs []int, target int) (int, bool) {
	left := 0
	right := len(inputs)

	mid := -1
	for left <= right {
		mid = (left + right) / 2
		if inputs[mid] == target {
			return mid, true
		}
		if target > inputs[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return mid, false

}
