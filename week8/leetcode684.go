package week8

// 使用并查集
// 时间复杂度 O(m) m 为边数
// 空间复杂度 O(n) n 为结点数

func findRedundantConnection(edges [][]int) []int {
	n := 0
	for _, edge := range edges {
		if edge[0] > n {
			n = edge[0]
		}
		if edge[1] > n {
			n = edge[1]
		}
	}
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if fa[x] == x {
			return x
		}
		fa[x] = find(fa[x])
		return fa[x]
	}
	var ans []int
	for _, edge := range edges {
		x := find(edge[0])
		y := find(edge[1])
		if x == y {
			ans = edge
		} else {
			fa[x] = y
		}
	}
	return ans
}
