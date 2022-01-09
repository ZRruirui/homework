package week8

// 把岛屿进行并查集，水直接指向 -1
// 时间复杂度 O,n*m*log(n) 其中 m*m 为遍历图，log(n) 为并查集
// 空间复杂度 O(m,n)
func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	fa := make([]int, m*n)
	var find func(x int) int
	// 编号
	num := func(i, j int) int {
		return i*n + j
	}
	// 并查集初始化
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				fa[num(i, j)] = -1
			} else {
				fa[num(i, j)] = num(i, j)
			}
		}
	}
	find = func(x int) int {
		if x == -1 || fa[x] == x {
			return x
		}
		fa[x] = find(fa[x])
		return fa[x]
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '1' {
				fa[num(i, j)] = -1
				continue
			}
			t1 := num(i+1, j)
			t2 := num(i, j+1)
			if i+1 < m && grid[i+1][j] == '1' {
				x := find(num(i, j))
				y := find(t1)
				if x != y {
					fa[y] = x
				}
			}
			if j+1 < n && grid[i][j+1] == '1' {
				x := find(num(i, j))
				y := find(t2)
				if x != y {
					fa[y] = x
				}
			}
		}
	}
	ans := 0
	for i := range fa {
		if find(i) == i {
			ans += 1
		}
	}
	return ans
}
