package prefixsum

type NumMatrix struct {
	matrix [][]int
	prefix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	n := len(matrix)
	m := len(matrix[0])
	prefix := make([][]int, n)
	for row := 0; row < n; row++ {
		prefix[row] = make([]int, m)
		for col := 0; col < m; col++ {
			if col == 0 {
				prefix[row][col] = matrix[row][col]
			} else {
				prefix[row][col] = prefix[row][col-1] + matrix[row][col]
			}
		}
	}
	return NumMatrix{
		matrix: matrix,
		prefix: prefix,
	}
}

// The brute force way to do this, WHICH WE SHOULD ALWAYS CONSIDER TO KNOW WE UNDERSTAND THE PROBLEM
// is too loop over matrix and add up
// From this we can recognize we can use prefix sum
func (m *NumMatrix) SumRegionBruteForce(row1 int, col1 int, row2 int, col2 int) int {
	result := 0
	for row := row1; row <= row2; row++ {
		for col := col1; col <= col2; col++ {
			result += m.matrix[row][col]
		}
	}
	return result
}

// we still have to loop n (number of rows) times, so this is not the constant solution asked for, though
// it is fast enough to pass
func (m *NumMatrix) SumRegionSimplePrefix(row1 int, col1 int, row2 int, col2 int) int {
	result := 0
	for row := row1; row <= row2; row++ {
		result += m.prefix[row][col2]
		if col1 != 0 {
			result -= m.prefix[row][col1-1]
		}
	}
	return result
}

func (m *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	result := 0
	for row := row1; row <= row2; row++ {
		result += m.prefix[row][col2]
		if col1 != 0 {
			result -= m.prefix[row][col1-1]
		}
	}
	return result
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

/*
 303. Range Sum Query - Immutable
Solved
Easy
Topics
Companies
Given an integer array nums, handle multiple queries of the following type:

Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
Implement the NumArray class:

NumArray(int[] nums) Initializes the object with the integer array nums.
int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).


Example 1:

Input
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
Output
[null, 1, -1, -3]

Explanation
NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3


Constraints:

1 <= nums.length <= 104
-105 <= nums[i] <= 105
0 <= left <= right < nums.length
At most 104 calls will be made to sumRange.
*/

// This is the sterotypical prefix sum
type NumArray struct {
	// data isn't necessary to maintain
	data []int
	// prefix is Inclusive prefix sum
	prefix []int
}

// NOTE: Constructor in problem
func NewNumArray(nums []int) NumArray {
	// calculate inclusive prefix sum
	// prefix[i] = prefix[i-1] + nums[i]
	prefix := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			prefix[i] = nums[i]
		} else {
			prefix[i] = prefix[i-1] + nums[i]
		}
	}

	return NumArray{
		data:   nums,
		prefix: prefix,
	}
}

// To get subarray sum from an Inclusive prefix sum
// we need only take right value and subtract every prior to left
// prefix[right] - prefix[left-1]
// Ex: given input of [1, 2, 3, 4, 5] the inclusive prefix is [1, 3, 6, 10, 15]
// SumRange(1, 4) = prefix[4] - prefix[1-1] = 15 - 1 = 14
func (this *NumArray) SumRange(left int, right int) int {
	result := this.prefix[right]
	if left != 0 {
		result -= this.prefix[left-1]
	}
	return result
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
