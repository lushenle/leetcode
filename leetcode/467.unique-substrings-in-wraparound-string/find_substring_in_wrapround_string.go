package leetcode

func findSubstringInWraproundString(p string) int {
	var count [26]int
	var length int

	for i := 0; i < len(p); i++ {
		if 0 < i && (p[i-1]+1 == p[i] || p[i-1] == p[i]+25) {
			length++
		} else {
			length = 1
		}

		b := p[i] - 'a'
		count[b] = max(count[b], length)
	}

	var ans int
	for i := 0; i < 26; i++ {
		ans += count[i]
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
