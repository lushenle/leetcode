package leetcode

import "math"

func reverse(x int) int {
	var ans int
	for x != 0 {
		if ans < math.MinInt32/10 || ans > math.MaxInt32/10 {
			return 0
		}

		ans = ans*10 + x%10
		x /= 10
	}

	return ans
}

func reverse1(x int) int {
	sign := 1

	// 处理负数
	if x < 0 {
		sign = -1
		x *= -1
	}

	ans := 0
	for x > 0 {
		temp := x % 10
		ans = ans*10 + temp
		x /= 10
	}

	// 还原 x 的符号
	ans *= sign

	// 处理溢出
	if ans > math.MaxInt32 || ans < math.MinInt32 {
		return 0
	}

	return ans
}
