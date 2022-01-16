package week9

func isMatch(s string, p string) bool {
	arrs := []byte(" " + s)
	arrp := []byte(" " + p)
	f := make([][]bool, len(arrs))
	for i := range f {
		f[i] = make([]bool, len(arrp))
	}
	f[0][0] = true
	for j := 1; j < len(arrp); j++ {
		if arrp[j] == '*' {
			f[0][j] = true
		} else {
			break
		}
	}
	for i := 1; i < len(arrs); i++ {
		for j := 1; j < len(arrp); j++ {
			if arrs[i] == arrp[j] || arrp[j] == '?' {
				f[i][j] = f[i-1][j-1]
			} else if arrp[j] == '*' {
				f[i][j] = f[i-1][j] || f[i][j-1]
			}
		}
	}
	return f[len(s)][len(p)]
}
