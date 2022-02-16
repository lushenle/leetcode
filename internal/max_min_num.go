package internal

func min(nums ...int) int {
	ans := nums[0]
	for _, v := range nums {
		if v < ans {
			ans = v
		}
	}
	return ans
}

func max(nums ...int) int {
	ans := nums[0]
	for _, v := range nums {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func maxElement(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
	}

	return nums[len(nums)-1]
}
