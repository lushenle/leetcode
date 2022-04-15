package leetcode

func rob(nums []int) int {
	var a, b int
	for idx, val := range nums {
		if idx%2 == 0 {
			a = max(a+val, b)
		} else {
			b = max(a, b+val)
		}
	}

	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
