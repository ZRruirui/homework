package week9

// 利用 大小写字母间的ASCII差
func toLowerCase(s string) string {
	arr := []byte(s)
	d := 'a' - 'A'
	for i, c := range arr {
		if c >= 'A' && c <= 'Z' {
			arr[i] = byte(int32(c) + d)
		}
	}
	return string(arr)
}
