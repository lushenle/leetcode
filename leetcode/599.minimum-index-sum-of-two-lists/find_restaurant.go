package leetcode

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	var ans []string
	list1Index := make(map[string]int)
	minIndexSum := math.MaxInt32

	for index, value := range list1 {
		list1Index[value] = index
	}

	for index, value := range list2 {
		if mIndex, ok := list1Index[value]; ok {
			if index+mIndex < minIndexSum {
				minIndexSum = index + mIndex
				ans = []string{value}
			} else if index+mIndex == minIndexSum {
				ans = append(ans, value)
			}
		}
	}

	return ans
}
