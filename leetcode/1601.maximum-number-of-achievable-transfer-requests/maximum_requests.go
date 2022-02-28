package leetcode

func maximumRequests(n int, requests [][]int) int {
	ans := 0
	count := 0
	m := make([]int, n)
	var dfs func(m []int, count int, index int)
	dfs = func(m []int, count int, index int) {
		if check(m) {
			ans = max(ans, count)
		}

		for i := index; i < len(requests); i++ {
			m[requests[i][0]]--
			m[requests[i][1]]++
			count++
			dfs(m, count, i+1)
			m[requests[i][0]]++
			m[requests[i][1]]--
			count--
		}
	}

	dfs(m, count, 0)

	return ans
}

func check(m []int) bool {
	for i := 0; i < len(m); i++ {
		if m[i] != 0 {
			return false
		}
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
