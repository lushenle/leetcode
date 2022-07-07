package leetcode

import (
	"sort"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	// preparing
	sort.Strings(dictionary)
	sentenceSlice := strings.Split(sentence, " ")

	for i, s := range sentenceSlice {
		for _, w := range dictionary {
			// replace if match
			if strings.HasPrefix(s, w) {
				sentenceSlice[i] = w
				break
			}
		}
	}

	// slice to string
	return strings.Join(sentenceSlice, " ")
}
