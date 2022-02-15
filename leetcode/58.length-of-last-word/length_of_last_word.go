package leetcode

func lengthOfLastWord(s string) int {
	last := len(s) - 1
	for s[last] == ' ' {
		last--
	}

	first := last
	for first >= 0 && s[first] != ' ' {
		first--
	}

	return last - first
}
