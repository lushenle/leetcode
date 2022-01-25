package leetcode

//func removeDuplicates(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i] == nums[i+1] {
//			nums = append(nums[:i], nums[i+1:]...)
//			i--
//		}
//	}
//
//	return len(nums)
//}

//func removeDuplicates(nums []int) int {
//	n := len(nums)
//	if n == 0 {
//		return 0
//	}
//
//	slow := 1
//	for fast := 1; fast < n; fast++ {
//		if nums[fast] != nums[fast-1] {
//			nums[slow] = nums[fast]
//			slow++
//		}
//	}
//
//	return slow
//}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	fast, slow := 0, 0
	for fast < len(nums)-1 {
		for nums[slow] == nums[fast] {
			slow++
			if slow == len(nums) {
				return fast + 1
			}
		}
		nums[fast+1] = nums[slow]
		fast++
	}
	return fast + 1
}
