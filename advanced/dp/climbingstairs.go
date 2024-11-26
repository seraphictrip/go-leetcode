package dp

import "fmt"

/*
You are given an integer array cost where cost[i] is the cost of ith step on a staircase.
Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0, or the step with index 1.

Return the minimum cost to reach the top of the floor.


The brute force is we need to try every option,
the sub problem is cost[n] + min(minCost(cost, n+1), minCost(cost, n+2))
the starting positions are 0 and 1



Example 1:

//      			10									15
//				 /		\							/		\
//			10+15		10+20					+20			+0
//			/   \
//       +20	+0
Input: cost = [10,15,20]
Output: 15
Explanation: You will start at index 1.
- Pay 15 and climb two steps to reach the top.
The total cost is 15.
Example 2:

Input: cost = [1,100,1,1,1,100,1,1,100,1]
Output: 6
Explanation: You will start at index 0.
- Pay 1 and climb two steps to reach index 2.
- Pay 1 and climb two steps to reach index 4.
- Pay 1 and climb two steps to reach index 6.
- Pay 1 and climb one step to reach index 7.
- Pay 1 and climb two steps to reach index 9.
- Pay 1 and climb one step to reach the top.
The total cost is 6.


Constraints:

2 <= cost.length <= 1000
0 <= cost[i] <= 999
*/

// TODO: remember to make the decision tree
// we need to figure out how to go from decision tree to code
// this is not yet something you are good at, we ned PRACTICE
func MinCostClimbingStairs(cost []int) int {
	memo := make(map[int]int, len(cost))
	return min(minCost(cost, 0, memo), minCost(cost, 1, memo))
}

func minCost(cost []int, n int, memo map[int]int) int {
	if n >= len(cost) {
		return 0
	}
	if cached, ok := memo[n]; ok {
		fmt.Println("CACHE: ", memo[n], "csize:", len(memo), "size:", len(cost))
		return cached
	}
	memo[n] = cost[n] + min(minCost(cost, n+1, memo), minCost(cost, n+2, memo))
	return memo[n]
}
