package leetcode

import (
	"strconv"
)

//func factor(m, n int) int {
//	if n == 0 {
//		return m
//	}
//
//	return factor(n, m%n)
//}

func factor(m, n int) int {
	for m != 0 {
		m, n = n%m, m
	}

	return n
}

//func removeDuplicates(s []string) []string {
//	e := map[string]bool{}
//	var r []string
//
//	for v := range s {
//		if e[s[v]] == false {
//			e[s[v]] = true
//			r = append(r, s[v])
//		}
//	}
//
//	return r
//}

//func simplifiedFractions(n int) []string {
//	var ans []string
//
//	if n == 0 || n == 1 {
//		return ans
//	}
//
//	for i := 2; i <= n; i++ {
//		for j := 1; j < i; j++ {
//			fac := factor(i, j)
//			ans = append(ans, fmt.Sprint(j/fac)+"/"+fmt.Sprint(i/fac))
//		}
//	}
//
//	return removeDuplicates(ans)
//}

func simplifiedFractions(n int) []string {
	var ans []string

	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if factor(j, i) == 1 {
				ans = append(ans, strconv.Itoa(j)+"/"+strconv.Itoa(i))
			}
		}
	}

	return ans
}
