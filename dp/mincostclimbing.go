package dp

import (
	"cmp"
	"fmt"
	"slices"
)

/*
You are given an integer array cost where cost[i] is the cost of ith step on a staircase.
 Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0, or the step with index 1.

Return the minimum cost to reach the top of the floor.



Example 1:

Input: cost = [10,15,20]
minStartingAt(0, ...) = min(10+15, 10+15+20, 10+20)
minStartingAt(1, ...) = min(15, 15+20)
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

func MinCostClimbingStairs(cost []int) int {
	// we can start at either position 0 or 1
	// so try both
	return min(calcCost(0, cost), calcCost(1, cost))
}

func calcCost(i int, cost []int) int {
	// base case, we have reached the top
	// we always move up one or two steps, so we always move towards base case
	if i >= len(cost) {
		return 0
	}
	return min(cost[i]+calcCost(i+1, cost), cost[i]+calcCost(i+2, cost))
}

func MinCostClimbingStairsTopDown(cost []int) int {
	// we can start at either position 0 or 1
	// so try both
	memo := make(map[int]int, len(cost))

	result := min(calcCostMemo(0, cost, memo), calcCostMemo(1, cost, memo))
	fmt.Println(memo)
	return result
}

func calcCostMemo(i int, cost []int, memo map[int]int) int {
	// base case, we have reached the top
	// we always move up one or two steps, so we always move towards base case
	if i >= len(cost) {
		return 0
	}
	if cached, ok := memo[i]; ok {
		return cached
	}
	memo[i] = min(cost[i]+calcCostMemo(i+1, cost, memo), cost[i]+calcCostMemo(i+2, cost, memo))
	return memo[i]
}

/*
Laws of recursion
1. Recursive call is correct
2. all paths lead to base case
3. Base cases must be correct
*/

func IsPartition[T cmp.Ordered](set []T, P ...[]T) bool {
	// P does not contain the empty set
	for _, p := range P {
		if len(p) == 0 {
			return false
		}
	}
	// Union of all p = set
	slices.Sort(set)
	if !slices.Equal(set, Union(P...)) {
		return false
	}

	// The intersection of any two distinct sets in P is empty
	// compare each p pairwise and make sure disjoint
	return false
}

func Union[T cmp.Ordered](sets ...[]T) []T {
	result := make([]T, 0)
	for i := 0; i < len(sets); i++ {
		result = append(result, sets[i]...)
	}
	slices.Sort(result)
	return result
}
