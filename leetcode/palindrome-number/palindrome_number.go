package leetcode

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}

	if x < 0 || x%10 == 0 {
		return false
	}

	tmpNum := x
	reversedDigit := 0

	for tmpNum > 0 {
		reversedDigit = reversedDigit*10 + tmpNum%10
		tmpNum = tmpNum / 10
	}

	return reversedDigit == x
}
