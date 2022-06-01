package leetcode

func maxProfit(prices []int) int {
	minPrice, maxPrice := prices[0], 0
	for i := 1; i < len(prices); i++ {
		maxPrice = max(maxPrice, prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}

	return maxPrice
}

func abs(a int) int {
	b := a >> 63
	return (a ^ b) - b
}

func max(a int, b int) int { return (a + b + abs(a-b)) / 2 }
func min(a, b int) int     { return a + (b-a)>>31&(b-a) }
