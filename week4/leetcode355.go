package week4

/*
	思路：
		1、 使用全局计数器计算创建时间，因为time.unix只会取到秒，可能存在创建时间相同情况。
		2、 采用全局 map 存用户 ID 和用户对象的映射
		3、 使用头插法链表保存用户发布的文章，可以保证按时间顺序倒着排列
		4、 采用 map 存用户关注另一用户
		5、 将用户和关注的用户的文章列表采用之前学的合并k个有序链表进行取top10
			a. 采用大顶堆实现合并过程
*/

type Twitter struct {
	db map[int]*User
	c  int64
}

func (this *Twitter) GetUser(userID int) *User {
	if this.db[userID] == nil {
		this.db[userID] = &User{
			head:   &ArticleNode{},
			Flower: make(map[int]*User),
		}
	}
	return this.db[userID]
}

type User struct {
	head   *ArticleNode
	Flower map[int]*User
}

func (u *User) Follow(userID int, user *User) {
	u.Flower[userID] = user
}

func (u *User) Unfollow(userID int) {
	delete(u.Flower, userID)
}

type Article struct {
	ID       int
	CreateAt int64
}

func (a *Article) Less(o *Article) bool {
	return a.CreateAt < o.CreateAt
}

func Constructor() Twitter {
	return Twitter{
		db: make(map[int]*User),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	a := Article{
		ID:       tweetId,
		CreateAt: this.c,
	}
	this.c += 1
	currentUser := this.GetUser(userId)
	InsertFront(currentUser.head, &a)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	user := this.GetUser(userId)
	h := NewHeap()
	ars := make([]*ArticleNode, 0, len(user.Flower)+1)
	if user.head.Next != nil {
		ars = append(ars, user.head.Next)
	}
	for _, f := range user.Flower {
		if f.head.Next != nil {
			ars = append(ars, f.head.Next)
		}
	}
	for _, a := range ars {
		h.push(a)
	}

	ans := make([]int, 0, 10)
	for len(ans) < 10 && h.Size() < 10 && !h.Empty() {
		a := h.pop()
		ans = append(ans, a.Val.ID)
		if a.Next != nil {
			h.push(a.Next)
		}
	}
	for len(ans) < 10 && !h.Empty() {
		e := h.pop()
		ans = append(ans, e.Val.ID)
	}
	return ans
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	user := this.GetUser(followerId)
	user.Follow(followeeId, this.GetUser(followeeId))
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	this.GetUser(followerId).Unfollow(followeeId)
}

type ArticleNode struct {
	Val  *Article
	Next *ArticleNode
}

func (l *ArticleNode) Less(o *ArticleNode) bool {
	return l.Val.CreateAt < o.Val.CreateAt
}

func InsertFront(head *ArticleNode, val *Article) {
	node := &ArticleNode{
		Val:  val,
		Next: head.Next,
	}
	head.Next = node
}

type heap struct {
	bucket []*ArticleNode
}

func NewHeap() *heap {
	return &heap{
		bucket: make([]*ArticleNode, 1, 10),
	}
}

func (h *heap) Size() int {
	return len(h.bucket)
}

func (h *heap) Empty() bool {
	return len(h.bucket) <= 1
}

func (h *heap) pop() *ArticleNode {
	if h.Empty() {
		return nil
	}
	e := h.bucket[1]
	h.bucket[1] = h.bucket[len(h.bucket)-1]
	h.bucket = h.bucket[:len(h.bucket)-1]
	h.fixDown()
	return e
}

func (h *heap) fixDown() {
	n := 1
	for n*2 < len(h.bucket) {
		c := n * 2
		if n*2+1 < len(h.bucket) {
			r := n*2 + 1
			if h.bucket[c].Less(h.bucket[r]) {
				c = r
			}
		}
		if h.bucket[n].Less(h.bucket[c]) {
			h.bucket[n], h.bucket[c] = h.bucket[c], h.bucket[n]
			n = c
		} else {
			break
		}
	}
}

func (h *heap) push(a *ArticleNode) {
	h.bucket = append(h.bucket, a)
	h.fixUp()
}

func (h *heap) fixUp() {
	n := len(h.bucket) - 1
	for n > 1 {
		if h.bucket[n/2].Less(h.bucket[n]) {
			h.bucket[n/2], h.bucket[n] = h.bucket[n], h.bucket[n/2]
			n = n / 2
		} else {
			break
		}
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
