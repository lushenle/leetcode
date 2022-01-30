package leetcode

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	var ans []string
	m := map[string]int{}
	for _, s := range []string{s1, s2} {
		for _, word := range strings.Split(s, " ") {
			m[word]++
		}
	}
	for key := range m {
		if m[key] == 1 {
			ans = append(ans, key)
		}
	}
	return ans
}
