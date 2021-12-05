package homework

func findOrder(numCourses int, prerequisites [][]int) []int {
	choise := make(map[int]bool)
	courseMap := make(map[int][]int)
	ans := make([]int, 0, numCourses)
	for _, v := range prerequisites {
		courseMap[v[0]] = append(courseMap[v[0]], v[1])
	}
	n := numCourses
	var fn func(int)
	fn = func(numCourses int) {
		if numCourses == 0 {
			return
		}
		for i := 0; i < n; i++ {
			if choise[i] {
				continue
			}
			if CanChoise(choise, courseMap[i]) {
				ans = append(ans, i)
				choise[i] = true
				fn(numCourses - 1)
				return
			}
		}
	}
	fn(numCourses)
	if len(ans) != n {
		return []int{}
	}
	return ans
}

func CanChoise(choise map[int]bool, pre []int) bool {
	for _, v := range pre {
		if !choise[v] {
			return false
		}
	}
	return true
}
