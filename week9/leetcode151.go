package week9

func reverseWords(s string) string {
	arr := []byte(s)
	s1 := make([]byte, 0, len(arr))
	s2 := make([]string, 0)
	for _, c := range arr {
		if c != ' ' {
			s1 = append(s1, c)
		} else {
			if len(s1) > 0 {
				s2 = append(s2, string(s1))
				s1 = s1[:0]
			}
		}
	}
	if len(s1) > 0 {
		s2 = append(s2, string(s1))
	}
	ans := ""
	for i := len(s2) - 1; i > 0; i-- {
		ans = ans + s2[i] + " "
	}
	return ans + s2[0]
}
