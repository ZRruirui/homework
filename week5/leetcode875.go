package week5

func minEatingSpeed(piles []int, h int) int {
	sum := 0
	r := 0
	for _, v := range piles {
		sum += v
		r = max(v, r)
	}
	l := max(1, sum/h)
	for l < r {
		mid := (l + r) / 2
		if valied(piles, h, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func valied(piles []int, h, sep int) bool {
	c := 0
	for i := range piles {
		c += piles[i] / sep
		if piles[i]%sep > 0 {
			c += 1
		}
		if c > h {
			return false
		}
	}
	return c <= h
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
