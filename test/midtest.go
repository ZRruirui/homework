package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	for a := 0; a < n-3; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		currentTarget := target - nums[a]
		for d := n - 1; d > a+2; d-- {
			// 固定 a,d
			if d < n-1 && nums[d+1] == nums[d] {
				continue
			}
			currentTarget = currentTarget - nums[d]
			for b := a + 1; b < d-1; b++ {
				if b > a+1 && nums[b] == nums[b-1] {
					continue
				}
				currentTarget = currentTarget - nums[b]
				if search(nums[b+1:d], currentTarget) {
					ans = append(ans, []int{nums[a], nums[b], currentTarget, nums[d]})
				}
				currentTarget = currentTarget + nums[b]
			}
			currentTarget = currentTarget + nums[d]
		}
	}

	return ans
}

func search(a []int, target int) bool {
	l := 0
	r := len(a) - 1
	for l <= r {
		mid := (l + r) / 2
		if a[mid] == target {
			return true
		}
		if a[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func main() {
	a := []int{1, 2, 4, 5, 6}
	fmt.Println(search(a, 3))
}
