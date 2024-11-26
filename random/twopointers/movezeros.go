package twopointers

import "slices"

/**/

// Input: nums = [0,1,0,3,12]
// [0,1,0,3,12]
//  L,R
// [1,0,0,3,12]
//    L R
//    L   R
// [1,3,0,0, 12]
//      L    R
// [1,3, 12, 0, 0]
// Output: [1,3,12,0,0

// with swap
func moveZeroes(nums []int) {
	L := 0
	// move L to first zero
	for L < len(nums) && nums[L] != 0 {
		L++
	}
	for R := L + 1; R < len(nums); R++ {
		if nums[R] != 0 {
			swap(nums, L, R)
			L++
		}
	}
}

// Fill in all positions 0 - len(nonzero), then fill in rest with 0s
func moveZeroesFindAndFill(nums []int) {
	L := 0
	for R := 0; R < len(nums); R++ {
		if nums[R] != 0 {
			nums[L] = nums[R]
			L++
		}
	}
	for L < len(nums) {
		nums[L] = 0
		L++
	}
}

func moveZerosSort(nums []int) {
	slices.SortStableFunc(nums, func(a, b int) int {
		// propagate zeros to back
		if a == 0 {
			return 1
		}
		if b == 0 {
			return -1
		}
		return 0
	})
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
