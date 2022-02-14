package leetcode

func singleNonDuplicate(nums []int) int {
	l, m, r := 0, 0, len(nums)-1
	for l < r {
		m = (r-l)/2 + l
		if m%2 == 1 {
			m--
		}
		if nums[m] == nums[m+1] {
			l = m + 2
		} else {
			r = m
		}
	}
	return nums[r]
}

func singleNonDuplicate1(nums []int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2

		if mid-1 >= 0 && nums[mid] == nums[mid-1] {
			mid--
		} else if mid+1 < len(nums) && nums[mid] == nums[mid+1] {
		} else {
			return nums[mid]
		}

		if ((mid - low) % 2) == 1 {
			high = mid - 1
		} else {
			low = mid + 2
		}
	}

	return -1
}
