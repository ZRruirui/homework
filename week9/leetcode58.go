package week9

func lengthOfLastWord(s string) int {
	arr := []byte(s)
	ans := 0
	for j := len(arr) - 1; j >= 0; j-- {
		if arr[j] != ' ' {
			arr = arr[:j+1]
			break
		}
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != ' ' {
			ans++
		} else {
			break
		}
	}
	return ans
}
