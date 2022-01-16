package week9

func findAnagrams(s string, p string) []int {
	h := hash(p)
	j := 0
	arr := []byte(s)
	sh := [26]int{}
	ans := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if i-j >= len(p) {
			sh[arr[j]-'a']--
			j++
		}
		sh[arr[i]-'a'] += 1
		if sh == h {
			ans = append(ans, j)
		}
	}
	return ans
}

func hash(s string) [26]int {
	res := [26]int{}
	for _, c := range []byte(s) {
		res[c-'a']++
	}
	return res
}
