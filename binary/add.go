package binary

import (
	"math/bits"
)

/*
Given two binary strings a and b, return their sum as a binary string.



Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"


Constraints:

1 <= a.length, b.length <= 104
a and b consist only of '0' or '1' characters.
Each string does not contain leading zeros except for the zero itself.

|binary		|	decimal |
|-----------|-----------|
|  0		|	0		|
|  1		|	1		|
|  11		|	3		|
|  100		|	4		|
|  101		|	5		|
|  110		|	6		|
|  111		|	7		|
|  1000		|	8		|
|  1001		|	9		|
|  1010		|	10		|
|  1011		|	11		|
|  1100		|	12		|
|  1101		|	13		|
|  1110		|	14		|
|  1111		|	15		|
|  10000	|	16		|
*/

func reverseBits(num uint32) uint32 {
	return bits.Reverse32(num)
}
