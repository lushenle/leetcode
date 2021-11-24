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
		lastDigit := tmpNum % 10
		reversedDigit = reversedDigit*10 + lastDigit
		tmpNum = tmpNum / 10
	}

	if reversedDigit == x {
		return true
	} else {
		return false
	}
}
