package leetcode

func minCost(costs [][]int) int {
	red, blue, green := 0, 0, 0

	for _, cost := range costs {
		red, blue, green = min(blue, green)+cost[0], min(red, green)+cost[1], min(red, blue)+cost[2]
	}

	return min(red, blue, green)
}

func min(values ...int) int {
	res := values[0]

	for _, val := range values {
		if val < res {
			res = val
		}
	}

	return res
}
