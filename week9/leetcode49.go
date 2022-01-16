package week9

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int]int)
	ans := make([][]string, 0)
	for _, str := range strs {
		key := hash(str)
		v, ok := m[key]
		if !ok {
			m[key] = len(ans)
			ans = append(ans, []string{str})
		} else {
			ans[v] = append(ans[v], str)
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
