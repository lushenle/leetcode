package leetcode

//func twoSum(nums []int, target int) []int {
//	length := len(nums)
//
//	for i := 0; i < length; i++ {
//		for j := i + 1; j < length; j++ {
//			if nums[i]+nums[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//
//	return []int{}
//}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for k, v := range nums {
		if i, ok := m[target-v]; ok {
			return []int{i, k}
		} else {
			m[v] = k
		}
	}

	return nil
}
