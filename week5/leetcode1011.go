package week5

func shipWithinDays(weights []int, days int) int {
	sum := 0
	maxWeight := 0
	for _, weight := range weights {
		sum = sum + weight
		if weight > maxWeight {
			maxWeight = weight
		}
	}
	left := maxWeight
	right := sum
	search := func(l, r int) int {
		for l <= r {
			mid := (l + r) / 2
			t := 0
			day := 1
			for _, weight := range weights {
				if t+weight > mid {
					day = day + 1
					t = weight
				} else {
					t = t + weight
				}
				if day > days {
					l = mid + 1
					break
				}
			}
			if day <= days {
				r = mid - 1
			}
		}
		return l
	}
	return search(left, right)
}
