package week9

func firstUniqChar(s string) int {
	m := make(map[byte]int)
	arr := []byte(s)
	for _, c := range arr {
		m[c]++
	}
	for i, c := range arr {
		if m[c] == 1 {
			return i
		}
	}
	return -1
}
