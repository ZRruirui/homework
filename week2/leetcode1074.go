package week2

//// 参考题解做的
//
//// 获取一维数组的
//// 时间复杂度 O(n) n 为一维数组长
//// 空间复杂度 O(n)
// 因为题解数用的行的前缀和，理解以后用列的前缀和做了下
// 计算一维度数组的方法数共用的
//func subarraySum(nums []int, k int) (ans int) {
//	mp := map[int]int{0: 1} // 一纬前缀和
//	for i, pre := 0, 0; i < len(nums); i++ {
//		pre += nums[i]
//		if _, ok := mp[pre-k]; ok {
//			ans += mp[pre-k]
//		}
//		mp[pre]++
//	}
//	return
//}
//
//// 获取二维矩阵的
//// 时间复杂度 O(n^2*m) n 为二维数组行数， m 为为二维数组列数
//// 空间复杂度 O(m) m 二维数组列数， 需要存每一个行的前缀合
//func numSubmatrixSumTarget(matrix [][]int, target int) (ans int) {
//	for i := range matrix { // 枚举上边界
//		sum := make([]int, len(matrix[0]))
//		for _, row := range matrix[i:] { // 枚举下边界
//			for c, v := range row {
//				sum[c] += v // 更新每列的元素和
//			}
//			// sum 为一行的前缀和
//			ans += subarraySum(sum, target)
//		}
//	}
//	return
//}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m := len(matrix)
	n := len(matrix[0])
	ans := 0
	for top := 0; top < m; top++ {
		sum := make([]int, n)
		for bottom := top; bottom < m; bottom++ {
			for j := 0; j < n; j++ {
				sum[j] = sum[j] + matrix[bottom][j]
			}
			ans = ans + subarraySum(sum, target)
		}
	}

	return ans
}
