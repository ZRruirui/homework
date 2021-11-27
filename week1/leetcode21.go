package week1

type ListNode struct {
	Val  int
	Next *ListNode
}

// 时间复杂度 O(n)
// 空间复杂度 O(1)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pro := &ListNode{}
	p := pro
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	for l1 != nil {
		p.Next = l1
		l1 = l1.Next
		p = p.Next
	}
	for l2 != nil {
		p.Next = l2
		l2 = l2.Next
		p = p.Next
	}
	return pro.Next
}
