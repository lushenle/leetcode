package internal

func Abs(a int) int {
	b := a >> 63
	return (a ^ b) - b
}

func Max(a int, b int) int {
	return (a + b + Abs(a-b)) / 2
}

func Min(a, b int) int {
	return a + (b-a)>>31&(b-a)
}
