package backtracking

/*
Given an array nums of distinct integers, return all the possible
permutations
. You can return the answer in any order.



Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
1, permute([2,3])
2, permute(1,3)
3, permute(1,2)
Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]
Example 3:

Input: nums = [1]
Output: [[1]]


Constraints:

1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.
*/

//			[1,2,3]
//		[1]					[2]					[3]
//	[1,2] [1,3]			[2,1], [2,3]		[3,1]		[3,2]
//
// [1,2,3]  [1,3,2]	[2,1,3]	[2,3,1]			[3,1,2]		[3,2,1]
func Permute(nums []int) [][]int {
	result := make([][]int, 0, Factorial(len(nums)))

	var dfs func(int)
	dfs = func(i int) {
		if i >= len(nums) {
			a := make([]int, len(nums))
			copy(a, nums)
			result = append(result, a)
			return
		}

		for j := i; j < len(nums); j++ {
			swap(nums, i, j)
			dfs(i + 1)
			// backtrack
			swap(nums, i, j)
		}

	}
	dfs(0)
	return result
}

func Factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
