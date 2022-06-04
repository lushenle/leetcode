package leetcode

func consecutiveNumbersSum(n int) int {
	var ans int
	for k, t := 1, 2*n; k*k < t; k++ {
		if t%k == 0 && (t/k+1-k)%2 == 0 {
			ans++
		}
	}

	return ans
}
