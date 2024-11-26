package slidingwindow

import "math"

// Q: Given an array, return true if there are two elements within a window of size that are equal

func CloseDuplicatesBruteForce(arr []int, k int) bool {
	n := len(arr)
	if n < 2 {
		return false
	}
	for L := 0; L < n; L++ {
		for R := L + 1; R < min(n, L+k); R++ {
			if arr[L] == arr[R] {
				return true
			}
		}
	}
	return false
}

func CloseDuplicatesWithHash(arr []int, k int) bool {
	seen := make(map[int]bool, k)
	L := 0

	for R := range len(arr) {
		if R >= L+k {
			seen[arr[L]] = false
			L += 1
		}
		if seen[arr[R]] {
			return true
		}
		seen[arr[R]] = true
	}
	return false

}

// Q: find the length of the longest subarray, with the same value in each position
// [1, 2, 2, 3, 2, 3, 3, 3]
//
//	L/R
//
// L, R
func LongestSameSubArray(arr []int) int {
	// initialize max length, 0 or 1 is fine unless len(a) == 0
	maxlength := 0

	L := 0
	for R := 0; R < len(arr); R++ {
		if arr[R] == arr[L] {
			maxlength = max(maxlength, R-L+1)
		} else {
			L = R
		}
	}

	return maxlength
}

// Find the minimum length subarray, where the sum is greater than or equal to the target.  Assume all values are positive
// [2,3,1,2,4,3] 6

func ShortestSubarray(arr []int, target int) int {
	total, L := 0, 0
	length := math.MaxInt
	for R := 0; R < len(arr); R++ {
		total += arr[R]
		for total >= target {
			length = min(length, R-L+1)
			total -= arr[L]
			L++
		}
	}
	if length == math.MaxInt {
		return 0
	}
	return length
}
