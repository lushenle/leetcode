package leetcode

func lexicalOrder(n int) []int {
	ans := make([]int, 0, n)

	var dfs func(start int)
	dfs = func(start int) {
		if start > n {
			return
		}

		for i := 0; i < 10 && start+i <= n; i++ {
			if start == 1 && i == 9 {
				break
			}
			ans = append(ans, start+i)
			dfs((start + i) * 10)
		}
	}

	dfs(1)

	return ans
}
