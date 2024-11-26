package twopointers

/*
You are given an integer array height of length n. There are n vertical lines
 drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container,
such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Draw a square, Height * Width
Height = min(arr[L], arr[R]) // can't slant
Width = R-L



Example 1:


Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
Example 2:

Input: height = [1,1]
Output: 1


Constraints:

n == height.length
2 <= n <= 105
0 <= height[i] <= 104
*/

func MaxArea(arr []int) int {
	maxarea := 0
	L, R := 0, len(arr)-1
	for L < R {
		area := (R - L) * min(arr[L], arr[R])
		maxarea = max(area, maxarea)
		if arr[L] < arr[R] {
			L++
		} else {
			R--
		}
	}

	return maxarea
}

/*
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.



Example 1:


Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]

| MaxLeft 	|
| MaxRight  |
| min(ml, mr)|

Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
Example 2:
                 L         R		maxLeft =4, maxRight=5
Input: height = [4,2,0,3,2,5]
Output: 9


Constraints:

n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105

*/

// Intuition, keep max heigh on both sides
func Trap(height []int) int {
	trapped := 0
	L, R := 0, len(height)-1
	maxL, maxR := height[L], height[R]
	for L < R {
		minwall := min(maxL, maxR)
		if maxL <= maxR {
			L++
			if minwall > height[L] {
				trapped += minwall - height[L]
			} else {
				maxL = height[L]
			}
		} else {
			R--
			if minwall > height[R] {
				trapped += minwall - height[R]
			} else {
				maxR = height[R]
			}
		}

	}
	return trapped
}

// L  L L L L R		maxLeft =4, maxRight=5
// [4,2,0,3,2,5]
/*
	 |
|	 |
|  | |
|| |||
|| |||



*/
