package arrays

// Insert val into arr at teh next open position
// Length is teh number of 'real' values in arr, and capacity
// is the size of underlying array
// Return length of new array
// O(1) assuming there is capacity
// Static arrays don't grow
func InsertEnd(arr []int, val int, length int, capacity int) int {
	if length < capacity {
		arr[length] = val
		return length + 1
	}

	return length
}

// Remove from the last position in the array if the array
// is not empty
// return new length of array
// O(1) remove from last position
func RemoveEnd(arr []int, length int) int {
	if length > 0 {
		// overwrite
		arr[length-1] = 0
		// reduce len
		return length - 1
	}
	return length
}

// Insert val at index i after shifting elements
// Assuming i isa  valid index and arr is not full
// return length of new array
// O(n)
func InsertMiddle(arr []int, i, val, length int) {
	// shift
	for index := length - 1; index >= i; i-- {
		arr[index+1] = arr[index]
	}

	arr[i] = val
}

func RemoveMiddle(arr []int, i, length int) {

}
