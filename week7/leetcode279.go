package week7

// 状态转移： f[i] = min(f[i-j*j]) + 1
// 时间复杂度 O(n^3/2)
// 空间复杂度 O(n)
func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minc := n
		for j := 1; j*j <= i; j++ {
			minc = min(minc, f[i-j*j])
		}
		f[i] = minc + 1
	}
	return f[n]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
