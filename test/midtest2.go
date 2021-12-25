package main

func hasPath(maze [][]int, start []int, destination []int) bool {
	dx := [4]int{-1, 0, 0, 1}
	dy := [4]int{0, -1, 1, 0}
	m := len(maze)
	n := len(maze[0])
	startX := start[0]
	startY := start[1]
	ans := false
	mark := map[[2]int]bool{
		[2]int{startX, startY}: true,
	}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		sx := x
		sy := y
		for i := range dx {
			if ans == true {
				return
			}
			nx := x + dx[i]
			ny := y + dy[i]
			if nx < 0 || nx >= m || ny < 0 || ny >= n || maze[nx][ny] == 1 {
				continue
			}
			for !(nx < 0 || nx >= m || ny < 0 || ny >= n) && maze[nx][ny] != 1 {
				x = nx
				y = ny
				nx = x + dx[i]
				ny = y + dy[i]
			}
			if x == destination[0] && y == destination[1] {
				ans = true
				return
			}
			p := [2]int{x, y}
			if !mark[p] {
				mark[p] = true
				dfs(x, y)
			}
			x = sx
			y = sy
		}
	}
	dfs(startX, startY)
	return ans
}
