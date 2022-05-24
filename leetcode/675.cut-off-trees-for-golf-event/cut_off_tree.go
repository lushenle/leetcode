package leetcode

import "sort"

func cutOffTree(forest [][]int) int {
	m, n := len(forest), len(forest[0])
	var trees [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if forest[i][j] > 1 {
				trees = append(trees, []int{forest[i][j], i*n + j})
			}
		}
	}
	sort.Slice(trees, func(i, j int) bool {
		return trees[i][0] < trees[j][0]
	})
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	bfs := func(start, end int) int {
		queue, explored := [][]int{{start, 0}}, map[int]bool{}
		for len(queue) > 0 {
			first := queue[0]
			cur, cost := first[0], first[1]
			if cur == end {
				return cost
			}
			x, y := cur/n, cur%n
			queue = queue[1:]
			for _, d := range dirs {
				nx, ny := x+d[0], y+d[1]
				p := nx*n + ny
				if 0 <= nx && nx < m && 0 <= ny && ny < n && forest[nx][ny] > 0 && !explored[p] {
					explored[p] = true
					queue = append(queue, []int{p, cost + 1})
				}
			}
		}
		return -1
	}

	ans := bfs(0, trees[0][1])
	for i := 0; i < len(trees)-1; i++ {
		if res := bfs(trees[i][1], trees[i+1][1]); res == -1 {
			return res
		} else {
			ans += res
		}
	}

	return ans
}
