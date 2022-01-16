package week9

func reverseOnlyLetters(s string) string {
	arr := []byte(s)
	l := 0
	r := len(arr) - 1
	for l < r {
		for l < r && !isLetter(arr[l]) {
			l++
		}
		for l < r && !isLetter(arr[r]) {
			r--
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	return string(arr)
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
