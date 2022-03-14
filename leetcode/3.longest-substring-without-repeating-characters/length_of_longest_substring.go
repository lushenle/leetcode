package leetcode

func lengthOfLongestSubstring(s string) int {
	right, left, ans := 0, 0, 0
	indexes := make(map[byte]int, len(s))

	for left < len(s) {
		if index, ok := indexes[s[left]]; ok && index >= right {
			right = index + 1
		}
		indexes[s[left]] = left
		left++
		ans = max(ans, left-right)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
