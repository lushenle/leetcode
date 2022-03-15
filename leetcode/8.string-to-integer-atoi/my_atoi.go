package leetcode

func myAtoi(s string) int {
	var digits []int
	// maxInt 不能直接使用 math.MaxInt32, 是无符号，有符号会溢出
	maxInt, signAllowed, whitespacesAllowed, sign := 1<<31, true, true, 1

	for _, c := range s {
		if c == ' ' && whitespacesAllowed {
			continue
		}
		if signAllowed {
			if c == '+' {
				signAllowed = false
				whitespacesAllowed = false
				continue
			} else if c == '-' {
				sign = -1
				signAllowed = false
				whitespacesAllowed = false
				continue
			}
		}
		if c < '0' || c > '9' {
			break
		}

		whitespacesAllowed, signAllowed = false, false
		digits = append(digits, int(c-48))
	}

	var num int
	place, lastLeading0Index := 1, -1

	for i, d := range digits {
		if d == 0 {
			lastLeading0Index = i
		} else {
			break
		}
	}
	if lastLeading0Index > -1 {
		digits = digits[lastLeading0Index+1:]
	}

	var rtnMax int
	if sign > 0 {
		rtnMax = maxInt - 1
	} else {
		rtnMax = maxInt
	}

	digitsCount := len(digits)
	for i := digitsCount - 1; i >= 0; i-- {
		num += digits[i] * place
		place *= 10
		if digitsCount-i > 10 || num > rtnMax {
			return sign * rtnMax
		}
	}

	num *= sign
	return num
}
