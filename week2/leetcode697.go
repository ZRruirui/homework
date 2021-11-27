package week2

// 时间复杂度: O(n)
// 空间复杂: O(n) hash map
func findShortestSubArray(nums []int) int {
	m := make(map[int]*Point)
	max := 1
	for i, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = &Point{
				Start: i,
				End:   i,
				Count: 1,
			}
			continue
		}
		m[num].End = i
		m[num].Count += 1
		if m[num].Count > max {
			max = m[num].Count
		}
	}
	ans := len(nums)
	for _, p := range m {
		if p.Count == max {
			if p.End-p.Start+1 < ans {
				ans = p.End - p.Start + 1
			}
		}
	}
	return ans
}

type Point struct {
	Start int
	End   int
	Count int
}
