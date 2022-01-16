package week9

func longestCommonPrefix(strs []string) string {
	arrs := make([][]byte, len(strs))
	minLen := int(1e9)
	for i, str := range strs {
		arrs[i] = []byte(str)
		if len(arrs[i]) < minLen {
			minLen = len(arrs[i])
		}
	}
	i := 0
	for i < minLen {
		c := arrs[0][i]
		for j := 1; j < len(arrs); j++ {
			if arrs[j][i] != c {
				return string(arrs[j][:i])
			}
		}
		i++
	}
	return string(arrs[0][:i])
}
