package week1

// 自己用双向链表实现的
//type linkNode struct {
//	Val   int
//	Next  *linkNode
//	Front *linkNode
//}
//
//type MyCircularDeque struct {
//	head *linkNode
//	tail *linkNode
//	cap  int
//	size int
//}
//
//func Constructor(k int) MyCircularDeque {
//	head := &linkNode{}
//	tail := &linkNode{}
//	head.Next = tail
//	tail.Front = head
//	return MyCircularDeque{
//		head: head,
//		tail: tail,
//		size: 0,
//		cap:  k,
//	}
//}
//
//func (this *MyCircularDeque) InsertFront(value int) bool {
//	if this.IsFull() {
//		return false
//	}
//	newNode := &linkNode{Val: value, Next: this.head.Next, Front: this.head}
//	this.head.Next.Front = newNode
//	this.head.Next = newNode
//	this.size += 1
//	return true
//}
//
//func (this *MyCircularDeque) InsertLast(value int) bool {
//	if this.IsFull() {
//		return false
//	}
//	newNode := &linkNode{Val: value, Next: this.tail, Front: this.tail.Front}
//	this.tail.Front.Next = newNode
//	this.tail.Front = newNode
//	this.size += 1
//	return true
//}
//
//func (this *MyCircularDeque) DeleteFront() bool {
//	if this.IsEmpty() {
//		return false
//	}
//	delNode := this.head.Next
//	delNode.Next.Front = this.head
//	this.head.Next = delNode.Next
//	this.size -= 1
//	return true
//}
//
//func (this *MyCircularDeque) DeleteLast() bool {
//	if this.IsEmpty() {
//		return false
//	}
//	delNode := this.tail.Front
//	delNode.Front.Next = this.tail
//	this.tail.Front = delNode.Front
//	this.size -= 1
//	return true
//}
//
//func (this *MyCircularDeque) GetFront() int {
//	if this.IsEmpty() {
//		return -1
//	}
//	return this.head.Next.Val
//}
//
//func (this *MyCircularDeque) GetRear() int {
//	if this.IsEmpty() {
//		return -1
//	}
//	return this.tail.Front.Val
//}
//
//func (this *MyCircularDeque) IsEmpty() bool {
//	return this.size == 0
//}
//
//func (this *MyCircularDeque) IsFull() bool {
//	return this.size == this.cap
//}

// 参考提示用数组实现的
type MyCircularDeque struct {
	head  int
	tail  int
	items []int
	empty bool
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		head:  0,
		tail:  0,
		empty: true,
		items: make([]int, k),
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.empty = false
	this.items[this.head] = value
	this.head = ((this.head + len(this.items)) - 1) % len(this.items)
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.empty = false
	this.tail = (this.tail + 1) % len(this.items)
	this.items[this.tail] = value
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % len(this.items)
	this.empty = this.head == this.tail
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail = (this.tail + len(this.items) - 1) % len(this.items)
	this.empty = this.head == this.tail
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	o := (this.head + 1) % len(this.items)
	return this.items[o]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.items[this.tail]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.empty
}
func (this *MyCircularDeque) IsFull() bool {
	return !this.empty && this.tail == this.head
}
