package leetcode

func countMaxOrSubsets(nums []int) int {
	n, m := len(nums), 0
	for _, num := range nums {
		m |= num
	}

	var dfs func(idx, val int) int
	dfs = func(idx, val int) int {
		if idx == n {
			if val == m {
				return 1
			}
			return 0
		}
		if nVal := val | nums[idx]; nVal == m {
			return dfs(idx+1, val) + (1 << (n - 1 - idx))
		} else {
			return dfs(idx+1, val) + dfs(idx+1, nVal)
		}
	}

	return dfs(0, 0)
}
