// Problem: https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

package leetcode

//func search(nums []int, target int) int {
//	for k, v := range nums {
//		if v == target {
//			return k
//		}
//	}
//	return -1
//}

func search(nums []int, target int) int {
	// Constraints:
	// 	1 <= nums.length <= 5000
	//if len(nums) == 0 {
	//	return -1
	//}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[low] { // target in the previous section
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] { // target in the back section
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if nums[low] == nums[mid] {
				low++
			}
			if nums[high] == nums[mid] {
				high--
			}
		}
	}
	return -1
}
