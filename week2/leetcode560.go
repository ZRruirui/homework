package week2

// 时间复杂度 O(n)
// 空间复杂度 O(n) 用map装前缀和
func subarraySum(nums []int, k int) (ans int) {
	mp := map[int]int{0: 1}
	for i, pre := 0, 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := mp[pre-k]; ok {
			ans += mp[pre-k]
		}
		mp[pre]++
	}
	return
}
