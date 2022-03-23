package leetcode

// 1 <= k <= n <= 109
// root 1: 10, 11, 12, ...
//      2: 20, 21, 22, ...
func findKthNumber(n int, k int) int {
	var dfs func(leaves, root int) int
	dfs = func(leaves, root int) int {
		if leaves > n {
			return 0
		}
		if root > n {
			root = n
		}
		return root - leaves + 1 + dfs(leaves*10, root*10+9)
	}

	current := 1
	k--
	for k > 0 {
		counts := dfs(current, current)
		if counts <= k {
			k -= counts
			current++
		} else {
			k--
			current *= 10
		}
	}

	return current
}
