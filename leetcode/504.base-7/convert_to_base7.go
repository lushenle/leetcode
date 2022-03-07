package leetcode

import "strconv"

func convertToBase7(num int) string {
	return strconv.FormatInt(int64(num), 7)
}

func abs(a int) int {
	b := a >> 63
	return (a ^ b) - b
}

func convertToBase71(num int) string {
	if num == 0 {
		return "0"
	}

	flag := num < 0
	num = abs(num)
	var s string

	for num > 0 {
		s = strconv.Itoa(num%7) + s
		num /= 7
	}

	if flag {
		s = "-" + s
	}

	return s
}
