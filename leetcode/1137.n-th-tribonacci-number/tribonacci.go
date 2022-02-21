package leetcode

func tribonacci(n int) int {
	m := make(map[int]int, 40)
	m[0] = 0
	m[1] = 1
	m[2] = 1

	for i := 3; i <= 37; i++ {
		m[i] = m[i-3] + m[i-2] + m[i-1]
	}

	return m[n]
}

func tribonacci1(n int) int {
	if n < 2 {
		return n
	}
	ans, first, zero := 1, 1, 0

	for n > 2 {
		ans, first, zero = ans+first+zero, ans, first
		n--
	}

	return ans
}
