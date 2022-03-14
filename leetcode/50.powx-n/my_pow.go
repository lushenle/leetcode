package leetcode

func myPow(x float64, n int) float64 {
	if x == 1 || n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	temp := myPow(x, n/2)
	if n%2 == 0 {
		return temp * temp
	} else {
		if n > 0 {
			return x * temp * temp
		} else {
			return temp * temp / x
		}
	}
}
