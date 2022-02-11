package leetcode

import (
	"sort"
)

//func min(m, n int) int {
//	if m < n {
//		return m
//	}
//
//	return n
//}

func min(a, b int) int {
	return a + (b-a)>>31&(b-a)
}

func abs(n int) int {
	m := n >> 63
	return (n ^ m) - m
}

func minimumDifference(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	sort.Ints(nums)
	ans := nums[k-1] - nums[0]

	for i := k; i < len(nums); i++ {
		if ans > nums[i]-nums[i-k+1] {
			ans = nums[i] - nums[i-k+1]
		}
	}

	return ans
}

func minimumDifference1(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	sort.Ints(nums)

	// 0 <= nums[i] <= 10^5
	ans := 100001

	for i := 0; i+k-1 < len(nums); i++ {
		ans = min(ans, abs(nums[i]-nums[i+k-1]))
	}

	return ans
}

func minimumDifference2(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	sort.Ints(nums)

	// 0 <= nums[i] <= 10^5
	ans := 100001

	for i := k - 1; i < len(nums); i++ {
		ans = min(ans, nums[i]-nums[i-k+1])
	}

	return ans
}
