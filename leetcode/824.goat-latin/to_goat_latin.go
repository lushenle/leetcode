package leetcode

import (
	"strings"
)

func toGoatLatin(sentence string) string {
	mVowel := map[string]string{"a": "a", "e": "e", "i": "i", "o": "o", "u": "u",
		"A": "A", "E": "E", "I": "I", "O": "O", "U": "U"}
	var ans string

	for idx, word := range strings.Split(sentence, " ") {
		if _, ok := mVowel[word[:1]]; !ok {
			word = word[1:] + word[:1]
		}

		ans += word + "ma" + strings.Repeat("a", idx+1) + " "
	}

	return strings.TrimRight(ans, " ")
}
