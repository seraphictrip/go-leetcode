package binarysearch

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"testing"
)

var BinarySearchTests = []struct {
	inputs []int
	target int
	index  int
	found  bool
}{
	{
		Range(0, 10, 1),
		3,
		3,
		true,
	},
	{
		Range(0, 10, 1),
		0,
		0,
		true,
	},
	{
		Range(0, 10, 1),
		9,
		9,
		true,
	},
	{
		Range(-100, 200, 1),
		0,
		100,
		true,
	},
	{
		Range(-100, 200, 2),
		0,
		50,
		true,
	},
	{
		Range(-100, 200, 2),
		1,
		51,
		false,
	},
}

func Range(start, count, inc int) []int {
	result := make([]int, count)
	for i := range result {
		result[i] = start + (inc * i)
	}
	return result
}
func TestBinarySearch(t *testing.T) {
	for i, e := range BinarySearchTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			index, found := BinarySearch(e.inputs, e.target)

			if found != e.found || index != e.index {
				t.Fatalf("BinarySearch(..., %v) = (%v %v), want (%v, %v)", e.target, index, found, e.index, e.found)
			}
		})
	}
}

/*
You are given an m x n integer matrix matrix with the following two properties:

Each row is sorted in non-decreasing order.
The first integer of each row is greater than the last integer of the previous row.
Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.



Example 1:


Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true
Example 2:


Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false


Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-104 <= matrix[i][j], target <= 104

*/

// Time Complexity: O(m + log(m*n))
// Space Complexity: O(m)
func searchMatrix(matrix [][]int, target int) bool {
	// detect row using info from column[0]
	// this is O(m)
	col0 := make([]int, len(matrix))
	for i := range matrix {
		col0[i] = matrix[i][0]
	}
	index, found := slices.BinarySearch(col0, target)
	if found {
		return true
	}

	if index-1 < 0 {
		return false
	}

	_, found = slices.BinarySearch(matrix[index-1], target)
	return found
}

var searchMatrixTests = []struct {
	matrix   [][]int
	target   int
	expected bool
}{
	{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true},
	{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false},
}

func TestSearchMatrix(t *testing.T) {
	for i, e := range searchMatrixTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := searchMatrix2(e.matrix, e.target)
			if actual != e.expected {
				t.Fatalf("searchMatrix(%v, %v) = %v, want %v", e.matrix, e.target, actual, e.expected)
			}
		})
	}
}

func bsearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

// [[ 1, 3, 5, 7 ]
//
//	[10, 11, 16, 20]
//	[23, 30, 34, 60]]
func searchMatrix2(matrix [][]int, target int) bool {
	ROWS, COLS := len(matrix), len(matrix[0])
	// 3, 4
	top, bot := 0, ROWS-1
	row := -1
	// 0, 2
	for top <= bot {
		// 1. top=0, bot=2
		// 2. top=0, bot=0
		row = (top + bot) / 2
		// 1. row=1
		// 2. row=0
		if target < matrix[row][0] {
			// 1. 3 < 10 (bot = 0)
			bot = row - 1

		} else if target > matrix[row][COLS-1] {
			top = row + 1
		} else {
			// we found the row
			break
		}
	}

	// if row is in range
	if row >= 0 && row < ROWS {
		// binary search row
		_, ok := slices.BinarySearch(matrix[row], target)
		return ok
	}
	return false

}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

//  func guessNumber(n int) int {
// 	l := 1
// 	r := n
// 	for l <= r {
// 		mid := (l+r)/2
// 		g := guess(mid)
// 		if g == 0 {
// 			return mid
// 		}
// 		if g == -1 {
// 			r = mid-1
// 		} else {
// 			l = mid+1
// 		}
// 	}
// 	return -1
//  }

var guessNumberTests = []struct {
	upper int
	pick  int
}{
	/*
		16383
		8191
		4095
		2047
		1023
		511
		255
		127
		63
		31
		15
		7
		3
		1
		0
		STEPS:  15
		PASS
	*/
	// {math.MaxInt16, -1},
	/*
		32766
		16383
		8191
		4095
		2047
		1023
		511
		255
		127
		63
		31
		15
		7
		3
		1
		0
		STEPS:  15
	*/
	{32766, -1},
}

func TestGuessNumber(t *testing.T) {
	for i, e := range guessNumberTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			count := 0
			fmt.Println(e.upper)
			check := func(x int) int {
				count++
				fmt.Println(x)
				if e.pick < x {
					return -1
				}
				if e.pick > x {
					return 1
				}
				return 0
			}
			// [1, INCLUSIVE_UPPER_BOUND]
			actual := binarySearchRange(0, e.upper, check)
			if actual != e.pick {
				t.Fatalf("Guess(%v) = %v, want %v", e.upper, actual, e.pick)
			}
			fmt.Println("STEPS: ", count)
		})
	}
}

func binarySearchRange(low, high int, check func(val int) int) int {
	for low <= high {
		mid := (low + high) / 2
		cmp := check(mid)
		if cmp == 0 {
			return mid
		}
		if cmp < 0 {
			high = mid - 1
		} else if cmp > 0 {
			low = mid + 1
		}
	}
	return -1
}

// Binary search are paritioning by 3, and explore if Ternary search requires 4 buckets etc

/*
There are two sides to the equation.  The partitioner, and the requestor.
If it is the partitioners repsonsibility to have quickest access, they should just
use a map. But if it is the requesters responsibility, they should ask for the key.
That key is the partition function that parititions the data in the most efficient way
for the requestor to both 1. find element, or know element is not in set.

A partitioning scheme with a key to each item is equivalent to a key (and a bucket) for each
item in the domain.  This system of keys can be expressed as F(x) = x over the domain of x
and can be represented as map of keys to values

*/
// Domain (D): int16
// min - middle | middle - min
// fmt.Println(math.MinInt16, (math.MinInt16+math.MaxInt16)/2, math.MaxInt16)
// -32768-0 0-32767
// Where does zero go?  Just have to make arbitrary decision
// [-32768...0] [1...32767]
// What info could I leak to help? find a specific item?
// first choice I can cut out half, but then I can no longer leak info
// if we add a new paritition, what can we put in it?
// lets just start with the item we are looking for
//  > n <
// [-32768...n-1] [n] [n+1..32767]

var IS_WAY_LESS_THAN = -2
var IS_LESS_THAN = -1
var IS_EQUAL = 0
var IS_GREATER_THAN = 1
var IS_WAY_GREATER_THAN = 2

// D: 0-10
// [0] [1] [2] [3] [4] [5] [6] [7] [8] [9] [10]

func ConceptualBinary(nums []int, target int) int {
	return target
}

var ConceptualBinaryTests = []struct {
	domain []int
	index  int
}{
	{Range(0, 1000, 1), 0},
}

func TestConceptualBinarySearch(t *testing.T) {
	for i, e := range BinarySearchTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			index := ConceptualBinary(e.inputs, e.target)

			if index != e.index {
				t.Fatalf("BinarySearch(..., %v) = %v, want %v", e.target, index, e.index)
			}
		})
	}
}

var playTests = []struct {
	r        []int
	expected []int
}{
	{RealRange(0, 5, 1), []int{0, 1, 2, 3, 4, 5}},
	{RealRange(0, 10, 1), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{RealRange(0, 10, 2), []int{0, 2, 4, 6, 8, 10}},
	{RealRange(10, 1, 1), []int{}},
	{RealRange(0, -100, -1), []int{}},
}

func TestPlay(t *testing.T) {
	for i, e := range playTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if !slices.Equal(e.r, e.expected) {
				t.Fatalf("%v", e.r)
			}

		})
	}
}

func RealRange(start, end, inc int) []int {
	if inc == 0 {
		panic("does not converge")
	}
	rang := make([]int, 0)
	cur := start
	for cur <= end {
		// we should detect we are going right distance or not
		rang = append(rang, cur)
		cur += inc
	}
	return rang
}

/*
You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.



Example 1:

Input: n = 5, bad = 4
Output: 4
Explanation:
call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true
Then 4 is the first bad version.
Example 2:

Input: n = 1, bad = 1
Output: 1


Constraints:

1 <= bad <= n <= 231 - 1

*/

// [fasle, false, false, true, true, true, true, true]

var FirstBadVersionTests = []struct {
	versions int
	bad      int
}{
	// [false, false, false, true, true]
	{5, 4},
	{5, 2},
	{1000, 6},
}

func TestFirstBadVersion(t *testing.T) {
	for i, e := range FirstBadVersionTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			check := func(n int) bool {
				return n >= e.bad
			}
			actual := firstBadVersion(e.versions, check)
			if actual != e.bad {
				t.Fatalf("firstBadVersion(%v) = %v, want %v", e.versions, actual, e.bad)
			}
		})
	}
}

// [false, true, true,  true, true]
func firstBadVersion(n int, isBadVersion func(n int) bool) int {
	low := 1
	high := n
	for low != high {
		mid := (low + high) / 2
		if isBadVersion(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

/*
Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer k such that she can eat all the bananas within h hours.



Example 1:

Input: piles = [3,6,7,11], h = 8
sum(piles) = 27
Output: 4
Example 2:

Input: piles = [30,11,23,4,20], h = 5
Output: 30
Example 3:

Input: piles = [30,11,23,4,20], h = 6
Output: 23


Constraints:

1 <= piles.length <= 104
piles.length <= h <= 109
1 <= piles[i] <= 109
*/

// k = sum(min)

func max(nums []int) int {
	m := math.MinInt
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	return m
}

func sum(nums []int) int {
	result := 0
	for _, n := range nums {
		result += n
	}
	return result
}

func minEatingSpeed(piles []int, h int) int {
	right := max(piles)
	left := 1

	res := right

	for left <= right {
		k := (left + right) / 2
		hours := 0
		for _, p := range piles {
			hours += int(math.Ceil(float64(p) / float64(k)))
		}
		if hours <= h {
			res = k
			right = k - 1
		} else {
			left = k + 1
		}

	}

	return res

}

//	1 <= k <= max(piles)
//
// sum(piles)/h <= k <= max(piles) * (len(n)/h)
func minEatingSpeed2(piles []int, h int) int {
	n := len(piles)
	m := max(piles)

	total := sum(piles)
	lowerbound := total / h
	scale := float64(n) / float64(h)
	upperbound := int(math.Ceil(float64(m) * scale))
	fmt.Printf("total = %v = Sum(%v)\n", total, piles)
	fmt.Printf("lowerbound = %v = %v/%v\n", lowerbound, total, h)
	fmt.Printf("scale = %v = %v/%v\n", scale, n, h)
	fmt.Printf("upperbound = %v = Ceil(%v*%v)\n", upperbound, m, scale)
	fmt.Println()

	left := lowerbound
	right := upperbound
	res := upperbound

	for left <= right {
		k := (left + right) / 2
		hours := 0
		for _, p := range piles {
			hours += int(math.Ceil(float64(p) / float64(k)))
		}
		if hours <= h {
			res = k
			right = k - 1
		} else {
			left = k + 1
		}

	}

	return res
}

var minEatingSpeedTests = []struct {
	piles    []int
	h        int
	expected int
}{
	{[]int{3, 6, 7, 11}, 8, 4},
	{[]int{30, 11, 23, 4, 20}, 5, 30},
}

func TestMinEatingSpeed(t *testing.T) {
	for i, e := range minEatingSpeedTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := minEatingSpeed2(e.piles, e.h)
			if actual != e.expected {
				t.Fatalf("minEatingSpeed(%v, %v) = %v, want %v", e.piles, e.h, actual, e.expected)
			}
		})
	}
}
