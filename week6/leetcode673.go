package week6

// 以 i 为右端的最长子段 f[i] = max(f[i-1,0]+1,1)
// 时间复杂度 O(n^2)
// 空间复杂度 O(n)
func findNumberOfLIS(nums []int) int {
	f := make([]int, len(nums))
	c := make([]int, len(nums))
	f[0] = 1
	c[0] = 1
	for i := 1; i < len(nums); i++ {
		// 以 i 为尾的最长递增子段的长度
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		// 以 i 为右端的最长子段的排列方式
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] && f[j]+1 == f[i] {
				c[i] += c[j]
			}
		}
		// 如果 i 不可能为右端
		// 子段长为1
		// 排列方式为1
		if f[i] == 0 {
			f[i] = 1
			c[i] = 1
		}
	}
	// 子段最长是多长
	m := 0
	for _, v := range f {
		if v > m {
			m = v
		}
	}
	// 最长子段有几种排列方式
	ans := 0
	for i, v := range f {
		if m == v {
			ans = ans + c[i]
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
