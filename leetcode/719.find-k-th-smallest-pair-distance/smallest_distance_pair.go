package leetcode

import "sort"

func smallestDistancePair(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	low := nums[1] - nums[0]

	for i := 2; i < n; i++ {
		low = min(low, nums[i]-nums[i-1])
	}

	high := nums[n-1] - nums[0]
	// ans: [low, high]
	// count(nums, mid): <= mid
	for low < high {
		mid := low + (high-low)/2
		if count(nums, mid) < k {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func count(nums []int, mid int) int {
	n, ans, i, j := len(nums), 0, 0, 1
	for j < n {
		if nums[j]-nums[i] <= mid {
			ans += j - i
			j++
		} else {
			i++
		}
	}

	return ans
}
