package fullyjustify

import (
	"fmt"
	"strings"
)

/*
Given an array of strings words and a width maxWidth, format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.

You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.

Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line does not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.

For the last line of text, it should be left-justified, and no extra space is inserted between words.

Note:

A word is defined as a character sequence consisting of non-space characters only.
Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
The input array words contains at least one word.


Example 1:

Input: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
Output:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]
Example 2:

Input: words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
Output:
[
  "What   must   be",
  "acknowledgment  ",
  "shall be        "
]
Explanation: Note that the last line is "shall be    " instead of "shall     be", because the last line must be left-justified instead of fully-justified.
Note that the second line is also left-justified because it contains only one word.
Example 3:

Input: words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"], maxWidth = 20
Output:
[
  "Science  is  what we",
  "understand      well",
  "enough to explain to",
  "a  computer.  Art is",
  "everything  else  we",
  "do                  "
]


Constraints:

1 <= words.length <= 300
1 <= words[i].length <= 20
words[i] consists of only English letters and symbols.
1 <= maxWidth <= 100
words[i].length <= maxWidth
*/

func FullJustify(words []string, maxWidth int) []string {
	rows := ToRows(words, maxWidth)
	output := make([]string, len(rows))
	for i := 0; i < len(rows)-1; i++ {
		output[i] = JustifyBetween(rows[i], maxWidth)

	}
	output[len(rows)-1] = JustifyLeft(rows[len(rows)-1], maxWidth)

	for _, line := range output {
		fmt.Printf("%d: %q\n", len(line), line)
	}
	return output
}

func ToRows(words []string, maxWidth int) [][]string {
	rows := make([][]string, 0)

	chcount := len(words[0])
	L := 0

	for R := 1; R < len(words); R++ {
		if chcount+1+len(words[R]) > maxWidth {
			rows = append(rows, words[L:R])
			L = R
			chcount = len(words[R])
		} else {
			chcount += len(words[R]) + 1 // add one for space
		}
	}
	rows = append(rows, words[L:])
	return rows
}

func JustifyBetween(words []string, maxWidth int) string {
	// for single words same as justifyleft
	if len(words) == 1 {
		return JustifyLeft(words, maxWidth)
	}
	// we have an "insert" between each word
	// these will grow left to right to fill space
	// "this", " ", "is", " ", "an"
	inserts := make([][]byte, len(words)-1)
	for i := range inserts {
		inserts[i] = []byte{' '}
	}

	total := len(strings.Join(words, ""))
	needed := maxWidth - total - len(inserts)
	for i := 0; i < needed; i++ {
		inserts[i%len(inserts)] = append(inserts[i%len(inserts)], ' ')
	}
	builder := strings.Builder{}
	for i := range words {
		builder.WriteString(words[i])
		if i < len(inserts) {
			builder.WriteString(string(inserts[i]))
		}

	}
	return builder.String()
}

func JustifyLeft(words []string, maxWidth int) string {
	buffer := make([]byte, maxWidth)
	for i := 0; i < maxWidth; i++ {
		buffer[i] = ' '
	}
	copy(buffer, strings.Join(words, " "))
	return string(buffer)
}
