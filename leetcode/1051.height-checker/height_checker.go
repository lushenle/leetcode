package leetcode

import "sort"

func heightChecker(heights []int) int {
	var ans int
	originHeights := make([]int, len(heights))
	copy(originHeights, heights)
	sort.Ints(heights)

	for i := 0; i < len(heights); i++ {
		if heights[i] != originHeights[i] {
			ans++
		}
	}

	return ans
}
