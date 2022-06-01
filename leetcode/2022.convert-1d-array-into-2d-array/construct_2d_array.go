package leetcode

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return nil
	}

	var ans [][]int
	for i := 0; i < len(original); i += n {
		ans = append(ans, original[i:i+n])
	}

	return ans
}
