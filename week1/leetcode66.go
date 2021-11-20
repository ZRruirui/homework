package week1

// 自己想的
func plusOne(digits []int) []int {
	ans := make([]int, len(digits)+1)
	ans[len(digits)] = 1
	for i := len(digits) - 1; i >= 0; i-- {
		d := digits[i] + ans[i+1]
		if d >= 10 {
			ans[i] = 1
		}
		ans[i+1] = d % 10
	}
	if ans[0] != 0 {
		return ans
	}
	return ans[1:]
}

// 参考题解提示
func plusOne2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] = digits[i] + 1
			return digits
		}
		digits[i] = 0
	}
	ans := make([]int, len(digits)+1)
	ans[0] = 1
	return ans
}
