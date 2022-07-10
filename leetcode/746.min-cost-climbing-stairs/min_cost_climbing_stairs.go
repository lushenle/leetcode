package leetcode

func minCostClimbingStairs(cost []int) int {
	a, b := cost[0], cost[1]
	for _, c := range cost[2:] {
		a, b = b, min(a, b)+c
	}

	return min(a, b)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
