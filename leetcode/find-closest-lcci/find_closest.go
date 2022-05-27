package leetcode

func findClosest(words []string, word1, word2 string) int {
	const max = 666666
	ans, idx1, idx2 := max, max, -max

	for idx, word := range words {
		if word == word1 {
			idx1 = idx
		} else if word == word2 {
			idx2 = idx
		}
		ans = min(ans, abs(idx1-idx2))
	}

	return ans
}

func abs(a int) int {
	b := a >> 63
	return (a ^ b) - b
}

func min(a, b int) int {
	return a + (b-a)>>31&(b-a)
}
