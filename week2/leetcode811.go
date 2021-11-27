package week2

import (
	"fmt"
	"strconv"
	"strings"
)

// 时间复杂度 O(n*k) k 为域名中点数
// 空间复杂度 O(N) 保存 map 映射
func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	for _, cpdomain := range cpdomains {
		arr := strings.Split(cpdomain, " ")
		c, _ := strconv.Atoi(arr[0])
		domain := arr[1]
		m[domain] += c
		for {
			i := strings.IndexByte(domain, '.')
			if i == -1 {
				break
			}
			domain = domain[i+1:]
			m[domain] += c
		}
	}
	ans := make([]string, 0, len(m))
	for k, v := range m {
		ans = append(ans, fmt.Sprintf("%d %s", v, k))
	}
	return ans
}
