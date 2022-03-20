package leetcode

func networkBecomesIdle(edges [][]int, patience []int) int {
	length, graph := len(patience), map[int][]int{}

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	distance, queue := make([]int, length), [][]int{{0, 0}}
	distance[0] = -1

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, other := range graph[current[0]] {
			if distance[other] == 0 {
				distance[other] = current[1] + 1
				queue = append(queue, []int{other, current[1] + 1})
			}
		}
	}

	var ans int
	for i := 1; i < length; i++ {
		d, p := distance[i]*2, patience[i]
		ans = max(ans, (d-1)/p*p+d+1)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
