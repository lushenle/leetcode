package leetcode

func numberOfLines(widths []int, s string) []int {
	if len(s) == 0 {
		return []int{0, 0}
	}

	var lines, curr int
	for _, c := range s {
		if curr+widths[c-'a'] > 100 {
			lines++
			curr = widths[c-'a']
		} else {
			curr += widths[c-'a']
		}
	}

	return []int{lines + 1, curr}
}
