package leetcode

// f(0) = 1
// f(1) = 9 + f(0)
// f(2) = 9 * 9 + f(1)
// f(3) = 9 * 9 * 8 + f(2)
// f(4) = 9 * 9 * 8 * 7 + f(3)
// ...
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}

	ans := 9
	for i := 9; i > 10-n; i-- {
		ans *= i
	}

	return ans + countNumbersWithUniqueDigits(n-1)
}
