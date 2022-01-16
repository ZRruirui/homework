package week9

func isAnagram(s string, t string) bool {
	return hash(s) == hash(t)
}

func hash(s string) [26]int {
	res := [26]int{}
	for _, c := range []byte(s) {
		res[c-'a']++
	}
	return res
}
