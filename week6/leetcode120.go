package week6

// 递推公式： f[i][j] = min(f[i-1][j], f[i-1][j-1]) + triangle[i][j]
// 时间复杂度 O(N^2)
// 空间复杂度 O(N^2)
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	f := make([][]int, len(triangle))
	for i := range f {
		f[i] = make([]int, m)
	}
	f[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				f[i][j] = f[i-1][j] + triangle[i][j]
			} else if j == i {
				f[i][j] = f[i-1][j-1] + triangle[i][j]
			} else {
				f[i][j] = min(f[i-1][j], f[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	ans := f[m-1][0]
	for j := 1; j < m; j++ {
		ans = min(ans, f[m-1][j])
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
