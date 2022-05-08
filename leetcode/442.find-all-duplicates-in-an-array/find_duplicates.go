package leetcode

func findDuplicates(nums []int) []int {
	encountered := map[int]bool{}
	var ans []int

	for idx := range nums {
		if !encountered[nums[idx]] {
			encountered[nums[idx]] = true
		} else {
			ans = append(ans, nums[idx])
		}
	}

	return ans
}
