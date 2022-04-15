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
