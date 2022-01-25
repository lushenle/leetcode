package leetcode

import (
	"math"
)

func checkPerfectNumber(num int) bool {
	result := 1

	if num%2 != 0 || num == 0 {
		return false
	}

	s := int(math.Sqrt(float64(num)))
	for i := 2; i <= s; i++ {
		if num%i == 0 {
			if i == s && i*i == num {
				result += i
			} else {
				result += i + num/i
			}
		}
	}

	return result == num
}
