package rotate

import "slices"

func Rotate(matrix [][]int) {
	Transpose(matrix)
	for i := 0; i < len(matrix); i++ {
		slices.Reverse(matrix[i])
	}
}

func Transpose(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// (i, j) => (j, n-i-1)
// (j, n-1-i) => (n-1-i, n-j-1)
// (n-1-i, n-j-1) => (n-1-j, i)
func Rotate90(matrix [][]int) {
	n := len(matrix)
	clone := make([][]int, n)
	for i := range clone {
		clone[i] = slices.Clone(matrix[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[j][n-i-1] = clone[i][j]
		}
	}
}

// Rotate in groups of 4
/*
n=2
[
	[P1,P4],  (0,0)=(i,j) (0,1)=(j, n-1-i)
	[P2, P3]  (1,0)=(n-1-j, i) (1,1)=(n-1-i, n-1-j)
]
	P1 = (i,j) TOP_LEFT
	P2 = (n-j-1, i) BOTTOM_LEFT
	P3 = (n-i-1,n-j-1) BOTTOM_RIGHT
	P4 = (j, n-1-i) TOP_RIGHT

	n=3
	[
	[1,P1,3], (0,1)=(i,j)
	[P2,5,P4], (1,0)=(n-1-j, i)   (1,2)=(j,n-1-i)
	[7,P3,9]] (2, 1)=(n-1-i, n-1-j)
*/
func Rotate90Inplace(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			p4r, p4c := Clockwise90(i, j, n)
			p3r, p3c := Clockwise90(p4r, p4c, n)
			p2r, p2c := Clockwise90(p3r, p3c, n)
			tmp := matrix[i][j]
			matrix[i][j] = matrix[p2r][p2c]
			matrix[p2r][p2c] = matrix[p3r][p3c]
			matrix[p3r][p3c] = matrix[p4r][p4c]
			matrix[p4r][p4c] = tmp

		}
	}

}
func Clockwise90(i, j, n int) (int, int) {
	return j, n - 1 - i
}

// (0,0,3) => (2,0)
// (0,1,3) => (1,0)
// (0,2,3) => (0,0)
func CounterClockwise90(i, j, n int) (int, int) {
	return n - 1 - j, i
}

type Queue []int

func (q *Queue) Enqueue(val int) {
	*q = append(*q, val)
}

func (q *Queue) Dequeue() int {
	old := *q
	first := old[0]
	*q = old[1:]
	return first
}
