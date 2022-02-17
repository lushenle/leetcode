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

func prob(N int, K int, r int, c int, steps int, dp [][][]float64) float64 {
	x := []int{1, 2, -1, -2, 1, 2, -1, -2}
	y := []int{2, 1, 2, 1, -2, -1, -2, -1}
	if steps == K {
		return 1
	}
	if dp[r][c][steps] > 0 {
		return dp[r][c][steps]
	}
	val := 0.0
	for i := 0; i < 8; i++ {
		row := r + x[i]
		col := c + y[i]
		if row < 0 || row >= N || col < 0 || col >= N {
			continue
		}
		val += float64(0.125 * prob(N, K, row, col, steps+1, dp))
	}
	dp[r][c][steps] = val
	return val
}

func knightProbability1(N int, K int, r int, c int) float64 {
	dp := make([][][]float64, N)
	for i := 0; i < N; i++ {
		dp[i] = make([][]float64, N)
		for j := 0; j < N; j++ {
			dp[i][j] = make([]float64, 101)
		}
	}
	return prob(N, K, r, c, 0, dp)
}

var dirs = []struct{ i, j int }{{-2, -1}, {-2, 1}, {2, -1}, {2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}}

func knightProbability2(n, k, row, column int) float64 {
	per := make([][][]float64, k+1)
	for i := range per {
		per[i] = make([][]float64, n)
		for j := range per[i] {
			per[i][j] = make([]float64, n)
		}
	}

	return dfs(per, n, row, column, k)
}

func dfs(per [][][]float64, n, r, c, s int) float64 {
	if r < 0 || r >= n || c < 0 || c >= n || s < 0 {
		return 0
	}

	if s == 0 {
		return 1.0
	}
	if per[s][r][c] != 0 {
		return per[s][r][c]
	}
	for _, v := range dirs {
		nextR, nextC := r+v.i, c+v.j
		per[s][r][c] = per[s][r][c] + dfs(per, n, nextR, nextC, s-1)/8
	}
	return per[s][r][c]
}
