package knapsack

import (
	"fmt"
	"slices"
)

/*
You have planned some train traveling one year in advance. The days of the year in which you will travel are given as an integer array days. Each day is an integer from 1 to 365.

Train tickets are sold in three different ways:

a 1-day pass is sold for costs[0] dollars,
a 7-day pass is sold for costs[1] dollars, and
a 30-day pass is sold for costs[2] dollars.
The passes allow that many days of consecutive travel.

For example, if we get a 7-day pass on day 2, then we can travel for 7 days: 2, 3, 4, 5, 6, 7, and 8.
Return the minimum number of dollars you need to travel every day in the given list of days.



Example 1:

Input: days = [1,4,6,7,8,20], costs = [2,7,15]
Output: 11
Explanation: For example, here is one way to buy passes that lets you travel your travel plan:
On day 1, you bought a 1-day pass for costs[0] = $2, which covered day 1.
On day 3, you bought a 7-day pass for costs[1] = $7, which covered days 3, 4, ..., 9.
On day 20, you bought a 1-day pass for costs[0] = $2, which covered day 20.
In total, you spent $11 and covered all the days of your travel.
Example 2:

Input: days = [1,2,3,4,5,6,7,8,9,10,30,31], costs = [2,7,15]
Output: 17
Explanation: For example, here is one way to buy passes that lets you travel your travel plan:
On day 1, you bought a 30-day pass for costs[2] = $15 which covered days 1, 2, ..., 30.
On day 31, you bought a 1-day pass for costs[0] = $2 which covered day 31.
In total, you spent $17 and covered all the days of your travel.


Constraints:

1 <= days.length <= 365
1 <= days[i] <= 365
days is in strictly increasing order.
costs.length == 3
1 <= costs[i] <= 1000
*/

func MincostTickets(days []int, costs []int) int {
	var helper func(i int) int
	helper = func(i int) int {
		if i >= len(days) {
			return 0
		}
		one := costs[0] + helper(i+1)
		seven := costs[1]
		thirty := costs[2]

		sevendays, _ := slices.BinarySearch(days, days[i]+7)
		thirtydays, _ := slices.BinarySearch(days, days[i]+30)

		if sevendays < len(days) {
			seven += helper(sevendays)
		}
		if thirtydays < len(days) {
			thirty += helper(thirtydays)
		}

		return min(one, seven, thirty)

	}

	return helper(0)
}

func MincostTicketsMemo(days []int, costs []int) int {
	var helper func(i int) int
	memo := make([]int, len(days)+1)
	for i := 0; i <= len(days); i++ {
		memo[i] = -1
	}
	helper = func(i int) int {
		if i >= len(days) {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		one := costs[0] + helper(i+1)
		seven := costs[1]
		thirty := costs[2]

		// find position where next would be
		sevendays, _ := slices.BinarySearch(days, days[i]+7)
		thirtydays, _ := slices.BinarySearch(days, days[i]+30)

		if sevendays <= len(days) {
			seven += helper(sevendays)
		}
		if thirtydays <= len(days) {
			thirty += helper(thirtydays)
		}

		memo[i] = min(one, seven, thirty)
		return memo[i]

	}

	ans := helper(0)
	return ans
}

func MincostTicketsBottomUp(days []int, costs []int) int {
	var helper func(i int) int
	table := make([]int, len(days)+1)
	// ORDER:
	// i requires i+[1, 7, 30]
	// small before big
	helper = func(i int) int {
		if i >= len(days) {
			return 0
		}
		one := costs[0] + helper(i+1)
		seven := costs[1]
		thirty := costs[2]

		// find position where next would be
		sevendays, _ := slices.BinarySearch(days, days[i]+7)
		thirtydays, _ := slices.BinarySearch(days, days[i]+30)

		if sevendays <= len(days) {
			seven += table[sevendays]
		}
		if thirtydays <= len(days) {
			thirty += table[thirtydays]
		}

		table[i] = min(one, seven, thirty)
		return table[i]

	}

	for i := len(days); i >= 0; i-- {
		table[i] = helper(i)
	}
	fmt.Println(table)
	return table[0]
}
