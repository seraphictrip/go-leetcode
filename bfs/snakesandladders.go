package bfs

func SnakesAndLadders(board [][]int) int {
	n := len(board)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	diceRolls := 0 // equivalent to level
	var neighbors Queue = make([]int, 0, n*n)
	neighbors.Enqueue(1)

	for len(neighbors) != 0 {
		size := len(neighbors)
		for size > 0 {
			cur := neighbors.Dequeue()
			if cur == n*n {
				return diceRolls
			}
			for roll := 1; roll <= 6; roll++ {
				// queue up all neighbors
				if cur+roll <= n*n {
					row, col := GetCoordinate(cur+roll, n)
					if !visited[row][col] {
						// mark visited, we don't need to try this in other plays
						visited[row][col] = true
						if board[row][col] == -1 {
							neighbors.Enqueue(cur + roll)
						} else {
							// snake or ladder
							neighbors.Enqueue(board[row][col])
						}
					}
				}
			}
			size--
		}
		diceRolls++
	}
	return -1
}

/*
GetCoordinate(1, 3)
row = 3 - ((1-1)/3) - 1 = 2
col = (1-1) % 3 = 0
[7,8,9]
[6,5,4]
[1,2,3]
GetCoordinate(1, 3)
row = 3 - ((1-1)/3) - 1 = 2
col = (1-1) % 3 = 0
GetCoordinate(1, 5)
row = 3 - ((5-1)/3) - 1 = 1
col = 3 - 1 - ((5-1) % 3) = 1
*/
// Given square label and n of nxn matrix
// get coordinates (row, col)
func GetCoordinate(pos int, n int) (int, int) {
	row := n - ((pos - 1) / n) - 1
	col := (pos - 1) % n

	if row%2 == n%2 {
		col = n - 1 - col
	}

	return row, col
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
