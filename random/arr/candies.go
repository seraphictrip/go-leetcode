package arr

/*

There are n kids with candies. You are given an integer array candies, where each candies[i] represents the number of candies the ith kid has, and an integer extraCandies, denoting the number of extra candies that you have.

Return a boolean array result of length n, where result[i] is true if, after giving the ith kid all the extraCandies, they will have the greatest number of candies among all the kids, or false otherwise.

Note that multiple kids can have the greatest number of candies.



Example 1:

Input: candies = [2,3,5,1,3], extraCandies = 3
Output: [true,true,true,false,true]
Explanation: If you give all extraCandies to:
- Kid 1, they will have 2 + 3 = 5 candies, which is the greatest among the kids.
- Kid 2, they will have 3 + 3 = 6 candies, which is the greatest among the kids.
- Kid 3, they will have 5 + 3 = 8 candies, which is the greatest among the kids.
- Kid 4, they will have 1 + 3 = 4 candies, which is not the greatest among the kids.
- Kid 5, they will have 3 + 3 = 6 candies, which is the greatest among the kids.
Example 2:

Input: candies = [4,2,1,1,2], extraCandies = 1
Output: [true,false,false,false,false]
Explanation: There is only 1 extra candy.
Kid 1 will always have the greatest number of candies, even if a different kid is given the extra candy.
Example 3:

Input: candies = [12,1,12], extraCandies = 10
Output: [true,false,true]


Constraints:

n == candies.length
2 <= n <= 100
1 <= candies[i] <= 100
1 <= extraCandies <= 50
*/

func KidsWithCandies(candies []int, extraCandies int) []bool {
	m := 0
	for _, candies := range candies {
		if candies > m {
			m = candies
		}
	}
	return Map(candies, func(c int) bool {
		return c+extraCandies >= m
	})
}

func Map[T any, V any](arr []T, transform func(T) V) []V {
	result := make([]V, len(arr))
	for i := range arr {
		result[i] = transform(arr[i])
	}
	return result
}

func Filter[T any](arr []T, filter func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range arr {
		if filter(item) {
			result = append(result, item)
		}
	}
	return result
}

func Reduce[T any, V any](arr []T, reducer func(acc V, cur T) V, acc V) V {
	for _, item := range arr {
		acc = reducer(acc, item)
	}
	return acc
}

func Some[T any](arr []T, predicate func(T) bool) bool {
	for _, item := range arr {
		if predicate(item) {
			return true
		}
	}
	return false
}

func Every[T any](arr []T, predicate func(T) bool) bool {
	for _, item := range arr {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func Unique[T any](arr []T, equal func(T) bool) bool {
	count := 0
	for _, item := range arr {
		if equal(item) {
			count++
			if count > 1 {
				return false
			}
		}
	}
	return count == 0
}
