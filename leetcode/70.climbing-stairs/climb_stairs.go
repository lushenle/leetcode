package leetcode

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		temp := n1 + n2
		n1 = n2
		n2 = temp
	}

	return n2
}
