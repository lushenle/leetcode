package leetcode

func removeOuterParentheses(s string) string {
	var ans []byte
	length := len(s)

	for i, j, cur := 0, 0, 0; i < length; i = j {
		j++
		cur++
		for j < length && cur > 0 {
			if s[j] == '(' {
				cur++
			} else {
				cur--
			}
			ans = append(ans, s[j])
			j++
		}
		ans = ans[:len(ans)-1]
	}

	return string(ans)
}
