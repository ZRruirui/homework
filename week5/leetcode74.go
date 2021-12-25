package week5

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	arr := make([]int, 0, m)
	for i := 0; i < m; i++ {
		arr = append(arr, matrix[i][0])
	}
	i := findFistLow(arr, target)
	return find(matrix[i], target)
}

func find(nums []int, target int) bool {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func findFistLow(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		mid := (l + r + 1) / 2
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}
