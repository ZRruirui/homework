package week4

const O = 'O'
const X = 'X'

func solve(board [][]byte) {
	bx := []int{-1, 0, 0, 1}
	by := []int{0, -1, 1, 0}
	mark := make(map[[2]int]bool)
	m := len(board)
	n := len(board[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		for i := range bx {
			nx := x + bx[i]
			ny := x + by[i]
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}
			if board[nx][ny] == O && !mark[[2]int{nx, ny}] {
				mark[[2]int{nx, ny}] = true
				dfs(nx, ny)
			}
		}
	}
	for i := 0; i < m; i++ {
		if board[i][0] == O && !mark[[2]int{i, 0}] {
			dfs(i, 0)
		}
		if board[i][n-1] == O && !mark[[2]int{i, n - 1}] {
			dfs(i, n-1)
		}
	}
	for j := 1; j < n-1; j++ {
		if board[0][j] == O && !mark[[2]int{0, j}] {
			dfs(0, j)
		}
		if board[m-1][j] == O && !mark[[2]int{m - 1, j}] {
			dfs(m-1, j)
		}
	}
	for i := range board {
		for j := range board[i] {
			if !mark[[2]int{i, j}] {
				board[i][j] = X
			}
		}
	}
}
