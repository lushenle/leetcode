package leetcode

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	var ans [][]int

	sort.Ints(arr)
	minAbs := arr[1] - arr[0]

	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] < minAbs {
			minAbs = arr[i+1] - arr[i]
		}
	}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == minAbs {
			ans = append(ans, []int{arr[i], arr[i+1]})
		}
	}

	return ans
}
