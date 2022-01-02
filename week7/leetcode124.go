package week7

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// 递归遍历树：
//    答案可能在 左子树，右子树，根，左子树+根，右子树 + 根， 左子树 + 右子树 + 根 这几种情况
func maxPathSum(root *TreeNode) int {
	ans := -1000
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftVal := dfs(root.Left)
		rightVal := dfs(root.Right)
		v := max(root.Val+leftVal, root.Val+rightVal)
		v = max(v, root.Val)
		ans = max(ans, v)
		ans = max(ans, leftVal+rightVal+root.Val)
		return v
	}
	return max(ans, dfs(root))
}
