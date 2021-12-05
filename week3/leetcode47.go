package week3

import "sort"

// used 表示下标 i 是否被使用
func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	used := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if used[i] || i > 0 && !used[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			used[i] = true
			backtrack(idx + 1)
			used[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}
