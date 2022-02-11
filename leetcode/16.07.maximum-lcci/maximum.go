package leetcode

func abs(n int) int {
	m := n >> 63
	return (n ^ m) - m
}

func maximum(a int, b int) int {
	return (a + b + abs(a-b)) / 2
}

func min(a, b int) int {
	return a + (b-a)>>31&(b-a)
}
