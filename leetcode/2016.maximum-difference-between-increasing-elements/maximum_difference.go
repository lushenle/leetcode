package leetcode

func maximumDifference(nums []int) int {
	ans := -1
	for i, preMin := 1, nums[0]; i < len(nums); i++ {
		if nums[i] > preMin {
			ans = max(ans, nums[i]-preMin)
		} else {
			preMin = nums[i]
		}
	}
	return ans
}

func abs(a int) int {
	b := a >> 63
	return (a ^ b) - b
}

func max(a int, b int) int {
	return (a + b + abs(a-b)) / 2
}
