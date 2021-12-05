package week3

import (
	"testing"
)

func Conv(a []int) *ListNode {
	head := &ListNode{}
	p := head
	for _, v := range a {
		p.Next = &ListNode{
			Val: v,
		}
		p = p.Next
	}
	return head.Next
}

func TestMergeKLists(t *testing.T) {
	p := Conv([]int{1, 4, 5})
	q := Conv([]int{1, 3, 4})
	r := Conv([]int{2, 6})
	h := mergeKLists([]*ListNode{p, q, r})
	display(h)
}
