package twopointers

/*
You are given an array of positive integers price where price[i] denotes the price of the ith candy
 and a positive integer k.

The store sells baskets of k distinct candies. The tastiness of a candy basket is the smallest
 absolute difference of the prices of any two candies in the basket.

Return the maximum tastiness of a candy basket.



Example 1:

Input: price = [13,5,1,8,21,2], k = 3

[1,8,21]
[1,2,5,8,13,21]

Output: 8
Explanation: Choose the candies with the prices [13,5,21].
The tastiness of the candy basket is: min(|13 - 5|, |13 - 21|, |5 - 21|) = min(8, 8, 16) = 8.
It can be proven that 8 is the maximum tastiness that can be achieved.
Example 2:

Input: price = [1,3,1], k = 2
[0,2,0]
[2,0,2]
[0,2,0]
Output: 2
Explanation: Choose the candies with the prices [1,3].
The tastiness of the candy basket is: min(|1 - 3|) = min(2) = 2.
It can be proven that 2 is the maximum tastiness that can be achieved.
Example 3:

Input: price = [7,7,7,7], k = 2
Output: 0
Explanation: Choosing any two distinct candies from the candies we have will result in a tastiness of 0.


Constraints:

2 <= k <= price.length <= 105
1 <= price[i] <= 109
*/

// Observation, maximum of single candy is max-min, so to maximize for 2 candies would take most expensive and least expensive
//

func MaximumTastiness(price []int, k int) int {
	// calculate all tastiness
	return 0
}

func Tastiness(c1, c2 int) int {
	if c1 < c2 {
		c1, c2 = c2, c1
	}
	return c1 - c2
}

func Choose(n, k int) int {
	if k == 0 {
		return 1
	}
	return (n * Choose(n-1, k-1)) / k
}

type MaxHeap []int

// Len is the number of elements in the collection.
func (h MaxHeap) Len() int {
	return len(h)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//   - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//   - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (h MaxHeap) Less(i int, j int) bool {
	// max, so need > rather than less
	return h[i] > h[j]
}

// Swap swaps the elements with indexes i and j.
func (h MaxHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	topIndex := len(old) - 1
	top := old[:topIndex]
	*h = old[:topIndex]
	return top
}

// return index of target or -1
// undefined for unsorted array
func BinarySearch(arr []int, target int) int {
	L, R := 0, len(arr)-1

	for L <= R {
		// TODO: I always forget to add the L so get stuck
		// how do I make myself remember this?
		// I forget because I usually use L+R/2, but it risks overflow
		mid := L + (R-L)/2
		if arr[mid] == target {
			return mid
		}
		if target < arr[mid] {
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	return -1
}
