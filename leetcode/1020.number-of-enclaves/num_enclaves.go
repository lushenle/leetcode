package leetcode

func isInGrid(grid [][]int, x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

func dfsNumEnclaves(grid [][]int, x, y int) {
	if !isInGrid(grid, x, y) || grid[x][y] == 0 {
		return
	}

	var dir = [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	grid[x][y] = 0

	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		dfsNumEnclaves(grid, nx, ny)
	}
}

func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				if grid[i][j] == 1 {
					dfsNumEnclaves(grid, i, j)
				}
			}
		}
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans++
			}
		}
	}
	
	return ans
}
