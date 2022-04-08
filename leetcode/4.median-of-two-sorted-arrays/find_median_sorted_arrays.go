package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := combine(nums1, nums2)
	return medianOf(nums)
}

// combine merge two sorted arrays
func combine(mis, njs []int) []int {
	lenMis, i := len(mis), 0
	lenNjs, j := len(njs), 0
	res := make([]int, lenMis+lenNjs)

	for k := 0; k < lenMis+lenNjs; k++ {
		if i == lenMis || i < lenMis && j < lenNjs && mis[i] > njs[j] {
			res[k] = njs[j]
			j++
			continue
		}

		if j == lenNjs || i < lenMis && j < lenNjs && mis[i] <= njs[j] {
			res[k] = mis[i]
			i++
		}
	}

	return res
}

// medianOf returns the median num
func medianOf(nums []int) float64 {
	length := len(nums)

	// m + n is even
	if length%2 == 0 {
		return float64(nums[length/2]+nums[length/2-1]) / 2.0
	}

	// m + n is odd
	return float64(nums[length/2])
}
