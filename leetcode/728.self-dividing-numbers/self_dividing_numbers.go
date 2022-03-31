package leetcode

func selfDividingNumbers(left int, right int) []int {
	isDividing := func(n int) bool {
		for i := n; i > 0; i /= 10 {
			if i%10 == 0 || n%(i%10) != 0 {
				return false
			}
		}
		return true
	}

	ans := make([]int, 0, right-left)
	for i := left; i <= right; i++ {
		if isDividing(i) {
			ans = append(ans, i)
		}
	}

	return ans
}
