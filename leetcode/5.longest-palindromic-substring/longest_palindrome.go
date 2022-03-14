package leetcode

func longestPalindrome(s string) string {
	length := len(s)

	if length <= 1 {
		return s
	}

	left, right, pl, pr := 0, -1, 0, 0
	for left < length {
		for right+1 < length && s[left] == s[right+1] {
			right++
		}

		for left-1 >= 0 && right+1 < length && s[left-1] == s[right+1] {
			left--
			right++
		}

		if right-left > pr-pl {
			pl, pr = left, right
		}

		left = (left+right)/2 + 1
		right = left
	}

	return s[pl : pr+1]
}
