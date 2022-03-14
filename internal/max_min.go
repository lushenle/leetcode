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

func flip(a int) int {
	return a ^ 1
}

func sign(a int) int {
	// return flip(a >> 31 &1)
	a = a >> 31 & 1
	return a ^ 1
}

func max1(a, b int) int {
	c := sign(a - b)
	d := flip(c)

	return a*c + b*d
}
