package leetcode

func hasAlternatingBits(n int) bool {
	m := n ^ n>>1

	return m&(m+1) == 0
}
