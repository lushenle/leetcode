package leetcode

func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	for left, right, cur := 0, 0, 1; right < len(nums); right++ {
		cur *= nums[right]
		for left <= right && cur >= k {
			cur /= nums[left]
			left++
		}
		ans += right - left + 1
	}

	return
}
func numSubarrayProductLessThanK1(nums []int, k int) (ans int) {
	prod, i := 1, 0
	for j, num := range nums {
		prod *= num
		for ; i <= j && prod >= k; i++ {
			prod /= nums[i]
		}
		ans += j - i + 1
	}
	return
}
