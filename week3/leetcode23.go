package week3

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode = &ListNode{}
	l := make([]*ListNode, 0, len(lists))
	for _, node := range lists {
		if node != nil {
			l = append(l, node)
		}
	}
	mergeKList(head, l)
	return head.Next
}

func mergeKList(head *ListNode, lists []*ListNode) {
	if len(lists) <= 0 {
		return
	}
	i := min(lists)
	head.Next = lists[i]
	lists[i] = lists[i].Next
	if lists[i] == nil {
		for ; i < len(lists)-1; i++ {
			lists[i] = lists[i+1]
		}
		lists = lists[:len(lists)-1]
	}
	mergeKList(head.Next, lists)
}

func min(list []*ListNode) int {
	m := int(math.MaxInt32)
	ans := 0
	for i := 0; i < len(list); i++ {
		if list[i] != nil && list[i].Val < m {
			m = list[i].Val
			ans = i
		}
	}
	return ans
}

func display(p *ListNode) {
	fmt.Println("---------------------")
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
