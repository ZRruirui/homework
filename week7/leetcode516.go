package week7

// i 表示右边界 j 表示做边界
// 状态转移 f[i][j] = f[i-1][j+1] + 2 当 s[i]=s[j] 时
// 					 max(f[i-1][j],f[i][j+1]
// 时间复杂度 O(n^2)
// 空间复杂度 O(n^2)
func longestPalindromeSubseq(s string) int {
	n := len(s)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
		f[i][i] = 1
	}
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				f[i][j] = f[i-1][j+1] + 2
			} else {
				f[i][j] = max(f[i-1][j], f[i][j+1])
			}
		}
	}
	return f[n-1][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
