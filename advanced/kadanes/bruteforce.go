package kadanes

/*
Given an integer array nums, find the
subarray
 with the largest sum, and return its sum.



Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
Example 2:

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.


Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104


Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

*/

// Q: find a non-empty subarray with the largest sum
// ex: [4,-1,2,-7, 3, 4] 7
// calculate largest subarray from each position
// by exhaustively calculating
// take max
// O(n^2)
func BruteForce(arr []int) int {
	maxsum := arr[0]
	for i := 0; i < len(arr); i++ {
		curSum := 0
		for j := i; j < len(arr); j++ {
			curSum += arr[j]
			maxsum = max(maxsum, curSum)
		}
	}
	return maxsum
}

// intuition, we only need to know max sum of previous, and can start fresh if max of position < zero
// [4,-1,2,-7, 3, 4] 7
// [4, 3, 5, -2, 3, 7]
// O(n)
func Kadanes(arr []int) int {
	maxsum := arr[0]
	cursum := 0
	for i := 0; i < len(arr); i++ {
		// no need to consider negative max sum
		// NOTE: this does not mean don't consider negative numbers, but rather max at position
		cursum = max(cursum, 0)
		cursum += arr[i]
		maxsum = max(maxsum, cursum)

	}
	return maxsum
}

// Sliding window version, in cases we need to know index of start, end
func SlidingWindow(arr []int) (start int, end int) {
	maxSum := arr[0]
	curSum := 0
	maxL, maxR := 0, 0
	L := 0

	for R := 0; R < len(arr); R++ {
		// just like kadane's, we can reset if we went negative
		if curSum < 0 {
			curSum = 0
			L = R
		}
		curSum += arr[R]
		if curSum > maxSum {
			maxSum = curSum
			maxL = L
			maxR = R
		}
	}
	return maxL, maxR
}
