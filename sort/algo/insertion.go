package algo

// Sort an array inplace
func InsertionSort(arr []int) {
	// invariant everything to the left of partition is in  sorted order
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				swap(arr, j-1, j)
			} else {
				// if we didn't need to swap no need to check further
				break
			}
		}
	}
}
