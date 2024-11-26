package trie

import "fmt"

/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.



Example 1:


Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true
Example 2:


Input: board =
[["A","B","C","E"],
 ["S","F","C","S"],
 ["A","D","E","E"]], word = "SEE"
Output: true
Example 3:


Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false
*/

type coordinate [2]int

// 1. Get dimensions (for boundary checks)
// 2. Create set of visted nodes
// 3. For each starting position, use DFS to look for exists
// 4. DFS with backtracking
//   - BASE CASE: Found word, return true
//   - boundary checks (am I on the board)
//   - validity check (not visited, character we are interested in)
//
// * Mark visited
// * check right
// * check down
// * check left
// * check up
// * backtrack (remove from visited)
func Exists(board [][]byte, word string) bool {
	ROWS, COLS := len(board), len(board[0])
	visited := make(map[coordinate]bool)

	var dfs func(row, col, k int) bool

	dfs = func(row, col, k int) bool {
		if k == len(word) {
			// we found the word
			return true
		}
		// boundary check
		if row < 0 || col < 0 {
			return false
		}
		if row >= ROWS || col >= COLS {
			return false
		}
		// already visited
		if visited[[2]int{row, col}] {
			return false
		}
		// actually what we want
		if word[k] != board[row][col] {
			return false
		}

		// mark visited
		visited[[2]int{row, col}] = true

		// check right
		if dfs(row, col+1, k+1) {
			return true
		}

		// check down
		if dfs(row+1, col, k+1) {
			return true
		}

		// check left
		if dfs(row, col-1, k+1) {
			return true
		}

		// check up
		if dfs(row-1, col, k+1) {
			return true
		}

		// backtrack
		visited[[2]int{row, col}] = false
		return false
	}
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

/*
https://leetcode.com/problems/word-search-ii/description/
Given an m x n board of characters and a list of strings words, return all words on the board.

Each word must be constructed from letters of sequentially adjacent cells,
where adjacent cells are horizontally or vertically neighboring.
The same letter cell may not be used more than once in a word.



Example 1:


Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]
Example 2:


Input: board = [["a","b"],["c","d"]], words = ["abcb"]
Output: []


Constraints:

m == board.length
n == board[i].length
1 <= m, n <= 12
board[i][j] is a lowercase English letter.
1 <= words.length <= 3 * 104
1 <= words[i].length <= 10
words[i] consists of lowercase English letters.
All the strings of words are unique
*/

// ASSUMPTION: will us DFS to look for words on board
// 1. Get board dimensions
// 2. init visited (could modify board with dummy value as well)
// 3. For each word, search board for word using DFS
//   - maybe build map of char : [coord] so only have to start at valid
//
// 4. DFS
//   - BASE CASE: word completely matched, return true
//   - boundary check
//   - validity check (node not visited + right char)
//   - mark visited
//   - check right, down, left, up
//   - backtrack (unmark visited)

// ... time out on some of test cases
func FindWords(board [][]byte, words []string) []string {
	ROWS, COLS := len(board), len(board[0])
	visited := make(map[coordinate]bool)

	// dfs at row, col using p(osition) in word
	var dfs func(row, col int, word string, p int) bool

	dfs = func(row, col int, word string, p int) bool {
		if p == len(word) {
			// we already found all chars
			return true
		}
		// boundary check
		if (row < 0 || row == ROWS) || (col < 0 || col == COLS) {
			return false
		}
		coor := [2]int{row, col}
		// validitity check
		if visited[coor] || word[p] != board[row][col] {
			return false
		}

		visited[coor] = true
		// right
		if dfs(row, col+1, word, p+1) {
			return true
		}
		// down
		if dfs(row+1, col, word, p+1) {
			return true
		}
		// left
		if dfs(row, col-1, word, p+1) {
			return true
		}
		// up
		if dfs(row-1, col, word, p+1) {
			return true
		}
		visited[coor] = false

		return false
	}

	set := make(map[string]bool, len(words))
	for _, w := range words {
		for i := 0; i < ROWS; i++ {
			for j := 0; j < COLS; j++ {
				// we need to reset visited, as it only relates to one path
				visited = make(map[coordinate]bool)
				if dfs(i, j, w, 0) {
					// we found, so add to result set
					set[w] = true
				}
			}
		}
	}
	result := make([]string, 0, len(set))
	for key, _ := range set {
		result = append(result, key)
	}

	return result
}

func FindWords1(board [][]byte, words []string) []string {
	ROWS, COLS := len(board), len(board[0])
	visited := make(map[coordinate]bool)
	resultset := make(map[string]bool)

	prefixTree := NewTrieNode()
	for _, word := range words {
		prefixTree.Insert(word)
	}

	var dfs func(row, col int, node *TrieNode, acc string)

	dfs = func(row, col int, node *TrieNode, acc string) {
		// boundary check first
		coor := [2]int{row, col}
		if (row < 0 || row == ROWS) || (col < 0 || col == COLS) || visited[coor] {
			return
		}

		ch := board[row][col]
		acc = fmt.Sprint(acc, string(ch))
		node = node.children[ch]
		if node == nil {
			return
		}

		if node.isWord {
			resultset[acc] = true
		}

		visited[coor] = true

		// right
		dfs(row, col+1, node, acc)

		// down
		dfs(row+1, col, node, acc)

		// left
		dfs(row, col-1, node, acc)

		// up
		dfs(row-1, col, node, acc)

		visited[coor] = false
	}

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			dfs(i, j, prefixTree, "")
		}
	}

	result := make([]string, 0, len(resultset))
	for key := range resultset {
		result = append(result, key)
	}
	return result
}

// Trie class
type TrieNode struct {
	children map[byte]*TrieNode
	isWord   bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode),
	}
}

func (t *TrieNode) Insert(word string) {
	cur := t
	for i := range word {
		node, ok := cur.children[word[i]]
		if ok {
			cur = node
		} else {
			node := NewTrieNode()
			cur.children[word[i]] = node
			cur = node
		}
	}
	cur.isWord = true
}

func (t *TrieNode) Search(word string) bool {
	cur := t
	for i := range word {
		node, ok := cur.children[word[i]]
		if !ok {
			return false
		}
		cur = node
	}
	return cur.isWord
}

func (t *TrieNode) StartsWith(word string) bool {
	cur := t
	for i := range word {
		node, ok := cur.children[word[i]]
		if !ok {
			return false
		}
		cur = node
	}
	return true
}
