package leetcode

func knightProbability(n int, k int, row int, column int) float64 {
	dp := make([]map[[2]int]float64, k+1)
	for i := range dp {
		dp[i] = make(map[[2]int]float64)
	}
	return dynamicProgramming(dp, n, k, row, column)
}

func dynamicProgramming(dp []map[[2]int]float64, n, k, r, c int) float64 {
	if !valid(n, r, c) {
		return 0.0
	}

	// valid here, in chessboard, return 1.0
	if k == 0 {
		return 1.0
	}

	row := min(r, n-1-r)
	col := min(c, n-1-c)
	if row > col {
		row, col = col, row
	}

	// the chessboard is square, so we support row <= col
	if v, exist := dp[k][[2]int{row, col}]; exist {
		return v
	}

	probability := float64(0)
	probability += dynamicProgramming(dp, n, k-1, r-2, c-1) / 8
	probability += dynamicProgramming(dp, n, k-1, r-1, c-2) / 8
	probability += dynamicProgramming(dp, n, k-1, r+2, c+1) / 8
	probability += dynamicProgramming(dp, n, k-1, r+1, c+2) / 8
	probability += dynamicProgramming(dp, n, k-1, r+2, c-1) / 8
	probability += dynamicProgramming(dp, n, k-1, r+1, c-2) / 8
	probability += dynamicProgramming(dp, n, k-1, r-2, c+1) / 8
	probability += dynamicProgramming(dp, n, k-1, r-1, c+2) / 8
	dp[k][[2]int{row, col}] = probability

	return probability
}

func valid(size, row, column int) bool {
	if row < 0 || column < 0 || row >= size || column >= size {
		return false
	}

	return true
}

func min(values ...int) int {
	minValue := 2147483647 // 1<<31 - 1

	for _, v := range values {
		if v < minValue {
			minValue = v
		}
	}

	return minValue
}
