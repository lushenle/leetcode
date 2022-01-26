package leetcode

//func removeElement(nums []int, val int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//
//	ans := 0
//	for _, v := range nums {
//		if v != val {
//			nums[ans] = v
//			ans++
//		}
//	}
//	return ans
//}

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)
	if right == 0 {
		return 0
	}

	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}
