package backtracking

import "math"

/*
Given an integer array nums of unique elements, return all possible
subsets
 (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.



Example 1:

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
Example 2:

Input: nums = [0]
Output: [[],[0]]


Constraints:

1 <= nums.length <= 10
-10 <= nums[i] <= 10
All the numbers of nums are unique.
*/

// Power set of A, P(A) = { {x}, {y}, {z}, {x, y}, {y, z}, {x, z}, {x, y, z}, {} }
// In set theory, the power set (or power set) of a Set A is defined as the set of all subsets of the Set A including the Set itself and the null or empty set. It is denoted by P(A).
//  Basically, this set is the combination of all subsets including null set, of a given set.
// power set will have 2^n elements, 2 can be though about as Inclusion or Exclusion
// for example empty set implies Exclusion for all elements, and full set (A) implies inclusion of all elements

// Cardinality of Power Set
// Cardinality represents the total number of elements present in a set. In case of power set, the cardinality will be the list of number of subsets of a set. The number of elements of a power set is written as |P (A)|, where A is any set. If A has ‘n’ elements then the formula to find the number of subsets of a set in a power set is given by:
// |P(A)| = 2^n

func PowerSetCardinality[T any](set []T) int {
	n := len(set)
	return int(math.Pow(2, float64(n)))
}

// [1,2]
// Cardinality: 2^2 = 4
// Decision Tree: Right=INCLUDE, Left=EXCLUDE (swappable, but general idea)
// 			[]
//		[1]		[]
//	 [2]  [1] [2] []

func PowerSet[T any](set []T) [][]T {
	ps := make([][]T, 0, PowerSetCardinality(set))
	subset := []T{}

	var createSubset func(int)
	createSubset = func(i int) {
		if i == len(set) {
			result := make([]T, len(subset))
			copy(result, subset)
			ps = append(ps, result)
			return
		}

		// with exclusion
		createSubset(i + 1)

		// with inclusion
		subset = append(subset, set[i])
		createSubset(i + 1)

		//  backtrack
		subset = subset[:len(subset)-1]
	}
	createSubset(0)
	return ps
}

func subsets(nums []int) [][]int {
	n := len(nums)
	cardinality := int(math.Pow(2, float64(n)))
	result := make([][]int, 0, cardinality)
	subset := []int{}
	var dfs func(int)
	dfs = func(i int) {
		if i >= n {
			// add copy to result
			c := make([]int, len(subset))
			copy(c, subset)
			result = append(result, c)
			return
		}
		// inclusion
		subset = append(subset, nums[i])
		dfs(i + 1)

		// exclusion (backtrack)
		subset = subset[:len(subset)-1]
		dfs(i + 1)
	}
	dfs(0)

	return result
}

/*
A recursive algorithm is used to generate the power set P(S) of any finite set S.

The operation F (e, T) is defined as:

F (e, T) = { X ∪ {e} | X ∈ T }

This returns each of the set X in T that has the element x.

If Set S = { }, then P(S) = { { } } is returned.

If not, the following algorithm is followed.

If e is an element in Set S, T = S {e} such that S { e } forms the relative complement of the element e in set S, the power set is generated by the following algorithm:

P(S) = P(T) ∪ F ( e, P(T))

To conclude, if the set S is empty, then the only element in the power set will be the null set. If not, the power set will become the union of all the subsets containing the particular element and the subsets not containing the particular element.
*/
