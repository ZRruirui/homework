package week7

// 状态转移：f[i] = f[i] | (f[j] && nums[j] >= i-j) // j 从0到i-1
// 时间复杂度 O(n^2)
// 空间复杂度 O(n)
func canJump(nums []int) bool {
	n := len(nums)
	f := make([]bool, len(nums))
	f[0] = true
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			f[i] = f[j] && (nums[j] >= i-j)
			if f[i] {
				break
			}
		}
	}
	return f[n-1]
}

// 贪心：求最远能到什么地方
// 时间复杂度 O(n)
// 空间复杂度 O(1)
//func canJump(nums []int) bool {
//	maxl := 0
//	for i := range nums {
//		if maxl < i{
//			break
//		}
//		maxl = max(maxl, i+nums[i])
//	}
//	return maxl >= len(nums)-1
//}
//
//func max(x,y int) int{
//	if x > y{
//		return x
//	}
//	return y
//}
