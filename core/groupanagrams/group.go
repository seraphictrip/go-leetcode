package groupanagrams

import "slices"

/*
Given an array of strings strs, group the
anagrams
 together. You can return the answer in any order.



Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]

Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Explanation:

There is no string in strs that can be rearranged to form "bat".
The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.
Example 2:

Input: strs = [""]

Output: [[""]]

Example 3:

Input: strs = ["a"]

Output: [["a"]]



Constraints:

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters.
*/

func GroupAnagrams(strs []string) [][]string {
	acc := make(map[string][]string)
	for i := range strs {
		id := identity(strs[i])
		if acc[id] == nil {
			acc[id] = make([]string, 1)
			acc[id][0] = strs[i]
		} else {
			acc[id] = append(acc[id], strs[i])
		}
	}
	result := make([][]string, 0, len(acc))
	for _, val := range acc {
		result = append(result, val)
	}
	return result
}

func identity(s string) string {
	runes := []rune(s)
	slices.Sort(runes)
	return string(runes)
}
