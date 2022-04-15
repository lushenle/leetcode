package leetcode

func firstMissingPositive(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, val := range nums {
		m[val] = val
	}

	for idx := 1; idx < len(nums)+1; idx++ {
		if _, ok := m[idx]; !ok {
			return idx
		}
	}

	return len(nums) + 1
}

func firstMissingPositive1(nums []int) int {
	// nums[idx] == idx+1
	for i := 0; i < len(nums); i++ {
		for 0 <= nums[i]-1 && nums[i]-1 < len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	// 第一个不存在的 idx+1 就是答案
	for idx := range nums {
		if nums[idx] != idx+1 {
			return idx + 1
		}
	}

	return len(nums) + 1
}
