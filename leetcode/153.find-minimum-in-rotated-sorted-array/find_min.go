package leetcode

func findMin(nums []int) int {
	low, mid, high := 0, 0, len(nums)-1

	for low+1 < high {
		mid = low + (high-low)>>1

		if nums[mid] == nums[high] {
			high = mid
		} else if nums[mid] > nums[high] {
			low = mid
		} else {
			high = mid
		}
	}

	if nums[low] <= nums[high] {
		return nums[low]
	}

	return nums[high]
}
