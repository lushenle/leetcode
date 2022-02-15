package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m + n; m > 0 && n > 0; i-- {
		if nums1[m-1] <= nums2[n-1] {
			nums1[i-1] = nums2[n-1]
			n--
		} else {
			nums1[i-1] = nums1[m-1]
			m--
		}
	}

	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}
