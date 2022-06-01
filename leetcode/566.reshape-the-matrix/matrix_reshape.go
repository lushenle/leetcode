package leetcode

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 || len(mat)*len(mat[0]) != r*c ||
		len(mat) == r && len(mat[0]) == c {
		return mat
	}

	ans := make([][]int, r)
	count, col := 0, len(mat[0])
	for i := range ans {
		ans[i] = make([]int, c)

		for j := range ans[i] {
			ans[i][j] = mat[count/col][count%col]
			count++
		}
	}

	return ans
}
