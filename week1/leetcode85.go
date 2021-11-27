package week1

// 这个题没想明白，看的题解
//func maximalRectangle(matrix [][]byte) (ans int) {
//	if len(matrix) == 0 {
//		return
//	}
//	m, n := len(matrix), len(matrix[0])
//	left := make([][]int, m)
//	for i, row := range matrix {
//		left[i] = make([]int, n)
//		for j, v := range row {
//			if v == '0' {
//				continue
//			}
//			if j == 0 {
//				left[i][j] = 1
//			} else {
//				left[i][j] = left[i][j-1] + 1
//			}
//		}
//	}
//	for j := 0; j < n; j++ { // 对于每一列，使用基于柱状图的方法
//		up := make([]int, m)
//		down := make([]int, m)
//		stk := []int{}
//		for i, l := range left {
//			for len(stk) > 0 && left[stk[len(stk)-1]][j] >= l[j] {
//				stk = stk[:len(stk)-1]
//			}
//			up[i] = -1
//			if len(stk) > 0 {
//				up[i] = stk[len(stk)-1]
//			}
//			stk = append(stk, i)
//		}
//		stk = nil
//		for i := m - 1; i >= 0; i-- {
//			for len(stk) > 0 && left[stk[len(stk)-1]][j] >= left[i][j] {
//				stk = stk[:len(stk)-1]
//			}
//			down[i] = m
//			if len(stk) > 0 {
//				down[i] = stk[len(stk)-1]
//			}
//			stk = append(stk, i)
//		}
//		for i, l := range left {
//			Height := down[i] - up[i] - 1
//			area := Height * l[j]
//			ans = max(ans, area)
//		}
//	}
//	return
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

// 听完助教讲解后做的
// 时间复杂度 O(m*n) m 为数组行数 n 为数组列数
// 空间复杂度 O(n*m) 开了一个等大的高数组，然后每行开一个等大的栈
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	// 申请等大的高容器
	m := len(matrix)
	n := len(matrix[0])
	heights := make([][]int, m)
	for i := range heights {
		heights[i] = make([]int, n)
	}
	// 计算以每一个行为底圆柱形
	for j := range matrix[0] {
		d := 0
		if matrix[0][j] == '1' {
			d = 1
		}
		heights[0][j] = d
	}
	for i := 1; i < m; i++ {
		for j := range matrix[i] {
			if matrix[i][j] == '1' {
				heights[i][j] = 1 + heights[i-1][j]
			}
		}
	}
	maxAres := 0
	for i := range heights {
		ares := largestRectangleArea(heights[i])
		if ares > maxAres {
			maxAres = ares
		}
	}
	return maxAres
}

// 时间复杂度 O(n)
// 空间复杂度 O(n)
func largestRectangleArea(heights []int) int {
	s := NewRectangleStack()
	maxAres := 0
	for _, h := range heights {
		if s.Empty() {
			s.Put(&Rectangle{
				Width:  1,
				Height: h,
			})
			continue
		}
		topRect := s.Top()
		width := 0
		if h < topRect.Height {
			for !s.Empty() && h < topRect.Height {
				rect := s.Pop()
				ares := (width + rect.Width) * rect.Height
				if ares > maxAres {
					maxAres = ares
				}
				width = width + rect.Width
				topRect = s.Top()
			}
		}
		s.Put(&Rectangle{
			Width:  width + 1,
			Height: h,
		})
	}
	width := 0
	for !s.Empty() {
		rect := s.Pop()
		ares := (width + rect.Width) * rect.Height
		if ares > maxAres {
			maxAres = ares
		}
		width = width + rect.Width
	}
	return maxAres
}

type Rectangle struct {
	Width  int
	Height int
}

type RectangleStack struct {
	items []*Rectangle
}

func NewRectangleStack() *RectangleStack {
	return &RectangleStack{
		items: make([]*Rectangle, 0),
	}
}

func (s *RectangleStack) Pop() *Rectangle {
	if s.Empty() {
		return nil
	}
	end := len(s.items) - 1
	item := s.items[end]
	s.items = s.items[:end]
	return item
}

func (s *RectangleStack) Put(item *Rectangle) {
	s.items = append(s.items, item)
}

func (s *RectangleStack) Empty() bool {
	return len(s.items) == 0
}

func (s *RectangleStack) Top() *Rectangle {
	if s.Empty() {
		return nil
	}
	return s.items[len(s.items)-1]
}
