package week7

// 状态转移 f[i] = min(f[i], f[j]+1) // j 为从 0 到 i-1 间可以到达 i 的
// 时间复杂度 O(n^2)
// 空间复杂度 O(n)
func jump(nums []int) int {
	n := len(nums)
	f := make([]int, len(nums))
	for i := range f {
		f[i] = 1e9
	}
	f[0] = 0
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j] >= i-j {
				f[i] = min(f[i], f[j]+1)
			}
		}
	}
	return f[n-1]
}

// 贪心
//func jump(nums []int) int {
//	if len(nums) == 1{
//		return 0
//	}
//	n := len(nums)
//	f := make([]int, len(nums))
//	f[0] = nums[0]
//	for i:=1;i<n;i++{
//		f[i] = max(f[i-1], nums[i]+i)
//	}
//	d := 1
//	i := 0
//	for f[i]<n-1{
//		i = f[i]
//		d ++
//	}
//	return d
//}
//
//func max(x,y int)int{
//	if x > y{
//		return x
//	}
//	return y
//}
