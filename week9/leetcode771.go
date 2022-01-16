package week9

func numJewelsInStones(jewels string, stones string) int {
	m := make(map[byte]bool)
	for _, c := range []byte(jewels) {
		m[c] = true
	}
	ans := 0
	for _, c := range []byte(stones) {
		if m[c] {
			ans++
		}
	}
	return ans
}
