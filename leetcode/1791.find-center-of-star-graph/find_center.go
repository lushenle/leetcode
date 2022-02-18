package leetcode

func findCenter1(edges [][]int) int {
	var nums []int
	m := make(map[int]int)
	nums = append(nums, edges[0][0], edges[0][1], edges[1][0], edges[1][1])

	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if v == 2 {
			return k
		}
	}

	return 0
}

func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}

	return edges[0][1]
}
