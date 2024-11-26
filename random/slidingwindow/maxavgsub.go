package slidingwindow

/*
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
 Any answer with a calculation error less than 10-5 will be accepted.
 * floating point math
 * fixed window size (K)



Example 1:

Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
Example 2:

Input: nums = [5], k = 1
Output: 5.00000


Constraints:

n == nums.length
1 <= k <= n <= 105
-104 <= nums[i] <= 104

*/

func FindMaxAverage(nums []int, k int) float64 {
	maxavg := 0.0
	sum := 0.0
	L, R := 0, 0
	// calculate first fixed window size
	for R < k {
		sum += float64(nums[R])
		R++
	}
	maxavg = sum / float64(k)

	// calc avg SUM(window)/k

	// slide window and repeat
	L++
	for R < len(nums) {
		sum -= float64(nums[L-1])
		sum += float64(nums[R])
		maxavg = max(maxavg, sum/float64(k))
		L++
		R++
	}

	return maxavg
}
