package backtracking

import (
	"fmt"
	"math/bits"
)

/*
A binary watch has 4 LEDs on the top to represent the hours (0-11), and 6 LEDs on the bottom to represent the minutes (0-59). Each LED represents a zero or one, with the least significant bit on the right.

For example, the below binary watch reads "4:51".
0: 0000
1: 0001
2: 0010
3: 0011
4: 0100
5: 0101
6: 0110
7: 0111
8: 1000
9: 1001
10: 1010
11: 1011
000000

Given an integer turnedOn which represents the number of LEDs that are currently on (ignoring the PM), return all possible times the watch could represent. You may return the answer in any order.

The hour must not contain a leading zero.

For example, "01:00" is not valid. It should be "1:00".
The minute must consist of two digits and may contain a leading zero.

For example, "10:2" is not valid. It should be "10:02".


Example 1:

Input: turnedOn = 1
Output: ["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]
Example 2:

Input: turnedOn = 9
Output: []


Constraints:

0 <= turnedOn <= 10
*/

// Given an input
func ReadBinaryWatch(turnedOn int) []string {
	result := make([]string, 0)
	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			// if count of on bits in binary representation
			// == turned on, add to output
			if bits.OnesCount(uint(h))+bits.OnesCount(uint(m)) == turnedOn {
				result = append(result, fmt.Sprintf("%v:%02d", h, m))
			}
		}
	}

	return result
}