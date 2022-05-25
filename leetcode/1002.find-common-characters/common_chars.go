package leetcode

import (
	"math"
)

func commonChars(words []string) []string {
	cnt := [26]int{}
	for i := range cnt {
		cnt[i] = math.MaxUint16
	}

	cntInWord := [26]int{}
	for _, word := range words {
		for _, char := range []byte(word) {
			cntInWord[char-'a']++
		}
		for i := 0; i < 26; i++ {
			if cntInWord[i] < cnt[i] {
				cnt[i] = cntInWord[i]
			}
		}

		for i := range cntInWord {
			cntInWord[i] = 0
		}
	}

	ans := make([]string, 0)
	for i := 0; i < 26; i++ {
		for j := 0; j < cnt[i]; j++ {
			ans = append(ans, string(rune(i+'a')))
		}
	}

	return ans
}
