package leetcode

func minDeletionSize(strs []string) int {
	var ans int
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] < strs[j-1][i] {
				ans++
				break
			}
		}
	}

	return ans
}
