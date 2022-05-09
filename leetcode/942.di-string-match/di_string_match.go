package leetcode

func diStringMatch(s string) []int {
	left, right := 0, len(s)
	var ans []int
	for _, c := range s {
		if c == 'I' {
			ans = append(ans, left)
			left++
		} else {
			ans = append(ans, right)
			right--
		}
	}

	ans = append(ans, right)
	return ans
}
