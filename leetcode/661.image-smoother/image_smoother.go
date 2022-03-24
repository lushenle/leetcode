package leetcode

func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	preSum, ans := make([][]int, m+1), make([][]int, m)
	preSum[0] = make([]int, n+1)

	for i := 0; i < m; i++ {
		preSum[i+1] = make([]int, n+1)
		ans[i] = make([]int, n)

		for j := 0; j < n; j++ {
			preSum[i+1][j+1] = preSum[i+1][j] + preSum[i][j+1] - preSum[i][j] + img[i][j]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			di, ui, dj, uj := max(0, i-1), min(m, i+2), max(0, j-1), min(n, j+2)
			ans[i][j] = (preSum[ui][uj] - preSum[ui][dj] - preSum[di][uj] + preSum[di][dj]) / ((ui - di) * (uj - dj))
		}
	}

	return ans
}

func flip(a int) int {
	return a ^ 1
}

func sign(a int) int {
	return flip(a >> 31 & 1)
}

func max(a, b int) int {
	c := sign(a - b)
	d := flip(c)
	return a*c + b*d
}

func min(a, b int) int {
	return a + (b-a)>>31&(b-a)
}

func imageSmoother1(img [][]int) [][]int {
	res := make([][]int, len(img))
	for i := range img {
		res[i] = make([]int, len(img[0]))
	}
	for y := 0; y < len(img); y++ {
		for x := 0; x < len(img[0]); x++ {
			res[y][x] = smooth(x, y, img)
		}
	}
	return res
}
func smooth(x, y int, img [][]int) int {
	count, sum := 1, img[y][x]
	// Check bottom
	if y+1 < len(img) {
		sum += img[y+1][x]
		count++
	}
	// Check Top
	if y-1 >= 0 {
		sum += img[y-1][x]
		count++
	}
	// Check left
	if x-1 >= 0 {
		sum += img[y][x-1]
		count++
	}
	// Check Right
	if x+1 < len(img[y]) {
		sum += img[y][x+1]
		count++
	}
	// Top Left
	if y-1 >= 0 && x-1 >= 0 {
		sum += img[y-1][x-1]
		count++
	}
	// Top Right
	if y-1 >= 0 && x+1 < len(img[0]) {
		sum += img[y-1][x+1]
		count++
	}
	// Bottom Left
	if y+1 < len(img) && x-1 >= 0 {
		sum += img[y+1][x-1]
		count++
	}
	//Bottom Right
	if y+1 < len(img) && x+1 < len(img[0]) {
		sum += img[y+1][x+1]
		count++
	}
	return sum / count
}
