package week3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	curr := postorder[len(postorder)-1]
	rootIndex := Find(inorder, curr)
	leftNode := buildTree(inorder[:rootIndex], postorder[:rootIndex])
	rightNode := buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])
	root := &TreeNode{Val: curr, Left: leftNode, Right: rightNode}
	return root
}

func Find(l []int, v int) int {
	for i := range l {
		if l[i] == v {
			return i
		}
	}
	return 0
}
