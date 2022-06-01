package leetcode

import "sort"

func makesquare(matchsticks []int) bool {
	if len(matchsticks) < 4 {
		return false
	}

	if sumArray(matchsticks)%4 != 0 {
		return false
	}
	quarter := sumArray(matchsticks) / 4
	for _, matchsticksLength := range matchsticks {
		if quarter < matchsticksLength {
			return false
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))
	edges := make([]int, 4)
	return dfs(matchsticks, edges, 0, quarter)
}

func dfs(nums, edges []int, index, target int) bool {
	if index == len(nums) {
		if edges[0] == target &&
			edges[1] == target &&
			edges[2] == target {
			return true
		}
		return false
	}

	for i := 0; i < 4; i++ {
		if edges[i]+nums[index] > target {
			continue
		}
		edges[i] += nums[index]
		if dfs(nums, edges, index+1, target) {
			return true
		}
		edges[i] -= nums[index]
	}

	return false
}

func sumArray(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}
