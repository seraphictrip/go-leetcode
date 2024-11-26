package twopointers

import (
	"strings"
	"unicode"
)

func IsPalindrome(arr []int) bool {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

func IsPalindrome2(arr []int) bool {
	L, R := 0, len(arr)-1
	for L < R {
		if arr[L] != arr[R] {
			return false
		}
		L++
		R--
	}
	return true
}

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*/
func IsCleansedPalindrome(s string) bool {
	s = cleanse(s)
	L, R := 0, len(s)-1
	for L < R {
		if s[L] != s[R] {
			return false
		}
		L++
		R--
	}
	return true
}

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters
// and removing all non-alphanumeric characters, it reads the same forward and backward.
// Alphanumeric characters include letters and numbers
func cleanse(s string) string {
	bytes := make([]rune, 0, len(s))
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			bytes = append(bytes, unicode.ToLower(ch))
		}
	}
	return string(bytes)
}

func IsCleansedPalindromeWalk(s string) bool {
	L, R := 0, len(s)-1
	for L < R {
		if !IsAlphanumeric(rune(s[L])) {
			L++
			continue
		}
		if !IsAlphanumeric(rune(s[R])) {
			R--
			continue
		}
		if unicode.ToLower(rune(s[L])) != unicode.ToLower(rune(s[R])) {
			return false
		}
		L++
		R--
	}
	return true
}
func IsAlphanumeric(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch)
}

func IsPalindrome3(s string) bool {
	s = strings.ToLower(s)
	L, R := 0, len(s)-1

	for L < R {
		if !IsAlphanumeric(rune(s[L])) {
			L++
			continue
		}
		if !IsAlphanumeric(rune(s[R])) {
			R--
			continue
		}
		if s[L] != s[R] {
			return false
		}
		L++
		R--

	}
	return true
}

/*
Given an integer array nums sorted in non-decreasing order,
remove the duplicates in-place such that each unique element appears only once.
The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the unique elements in the order
they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
Return k.
Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
If all assertions pass, then your solution will be accepted.



Example 1:

Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
Example 2:

Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).


Constraints:

1 <= nums.length <= 3 * 104
-100 <= nums[i] <= 100
nums is sorted in non-decreasing order.

*/

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	L, R := 1, 1
	for R < len(nums) {
		if nums[R] != nums[L-1] {
			swap(nums, L, R)
			L++
		}
		R++
	}
	return L
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func RemoveDuplicatesConcept(nums []int) int {
	lastSeen := nums[0]
	i := 0
	filtered := filter(nums, func(num int) bool {
		if i == 0 {
			i++
			return true
		}
		if num != lastSeen {
			lastSeen = num
			return true
		}
		return false
	})

	copy(nums, filtered)

	return len(filtered)
}

func filter[T any](arr []T, fn func(T) bool) []T {
	result := make([]T, 0, len(arr))

	for _, item := range arr {
		if fn(item) {
			result = append(result, item)
		}
	}

	return result
}
