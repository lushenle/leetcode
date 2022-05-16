package leetcode

func isAlienSorted(words []string, order string) bool {
	if len(words) < 2 {
		return true
	}

	m := make(map[byte]int)

	for i := 0; i < len(order); i++ {
		m[order[i]] = i
	}

	for i := 0; i < len(words)-1; i++ {
		pointer, word, wordPlus := 0, words[i], words[i+1]
		for pointer < len(word) && pointer < len(wordPlus) {
			if m[word[pointer]] > m[wordPlus[pointer]] {
				return false
			}
			if m[word[pointer]] < m[wordPlus[pointer]] {
				break
			} else {
				pointer = pointer + 1
			}
		}
		if pointer < len(word) && pointer >= len(wordPlus) {
			return false
		}
	}

	return true
}
