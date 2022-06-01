package leetcode

func minOperations(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	var ans int
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			ans += nums[i-1] - nums[i] + 1
			nums[i] = nums[i-1] + 1
		}
	}

	return ans
}
