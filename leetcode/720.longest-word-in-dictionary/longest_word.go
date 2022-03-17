package leetcode

import "strings"

func longestWord(words []string) (ans string) {
	m := map[string]bool{}
	for _, word := range words {
		m[word] = true
	}

out:
	for _, word := range words {
		for i := 1; i < len(word); i++ {
			if !m[word[:i]] {
				continue out
			}
		}
		if l1, l2 := len(word), len(ans); l1 > l2 || (l1 == l2 && strings.Compare(word, ans) < 0) {
			ans = word
		}
	}

	return
}
