package leetcode

import "strconv"

func optimalDivision(nums []int) string {
	length := len(nums)
	if length == 1 {
		return strconv.Itoa(nums[0])
	}

	if length == 2 {
		return strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1])
	}

	ans := strconv.Itoa(nums[0]) + "/(" + strconv.Itoa(nums[1])
	for i := 2; i < length; i++ {
		ans = ans + "/" + strconv.Itoa(nums[i])
	}

	return ans + ")"
}
