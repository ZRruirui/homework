package week5

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == nums[r] {
			r--
		} else if nums[mid] < nums[r] {
			r = mid
		} else {
			l = l + 1
		}
	}
	return nums[l]
}
