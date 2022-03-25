package leetcode

func trailingZeroes(n int) int {
	return n/5 + n/25 + n/125 + n/625 + n/3125
}

func trailingZeroes1(n int) (ans int) {
	for n > 0 {
		n /= 5
		ans += n
	}
	return
}
