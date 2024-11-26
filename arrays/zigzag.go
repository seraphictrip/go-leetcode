package arrays

import (
	"fmt"
	"strings"
)

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);


Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I
Example 3:

Input: s = "A", numRows = 1
Output: "A"


Constraints:

1 <= s.length <= 1000
s consists of English letters (lower-case and upper-case), ',' and '.'.
1 <= numRows <= 1000
*/

func Convert(s string, numRows int) string {
	matrix := make([][]byte, numRows)
	for i := range matrix {
		matrix[i] = make([]byte, 0)
	}
	indexMatrix := make([][]int, numRows)
	for i := range indexMatrix {
		indexMatrix[i] = make([]int, 0)
	}

	dir := 1
	i := 0
	for ptr := 0; ptr < len(s); ptr++ {
		if i == 0 {
			dir = 1
		}
		if i == numRows-1 {
			dir = -1
		}
		matrix[i] = append(matrix[i], s[ptr])
		indexMatrix[i] = append(indexMatrix[i], ptr)
		i += dir

	}
	var result strings.Builder
	for i := 0; i < numRows; i++ {
		result.WriteString(string(matrix[i]))
	}
	fmt.Println(indexMatrix)
	return result.String()
}

// func row(i int, numRows int) int {
// 	if i < numRows {
// 		return i
// 	}
// }
