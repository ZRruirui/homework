package week6

// f[i] = f[i-1] + f[i-2] 这是一个斐波那契数列 只需保存最近两个数就可以了
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	a := 1
	b := 1
	ans := 0
	for i := 2; i <= n; i++ {
		ans = a + b
		a = b
		b = ans
	}
	return ans
}
