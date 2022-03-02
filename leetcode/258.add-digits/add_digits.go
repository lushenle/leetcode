package leetcode

func addDigits0(num int) int {
	return (num-1)%9 + 1
}

func addDigits1(num int) int {
	if num <= 9 {
		return num
	}

	for num > 9 {
		var ans int
		for num != 0 {
			ans += num % 10
			num /= 10
		}
		num = ans
	}

	return num
}

func addDigits(num int) int {
	if num <= 9 {
		return num
	}

	if num > 9 {
		num %= 9
		if num == 0 {
			return 9
		}
	}

	return num
}
