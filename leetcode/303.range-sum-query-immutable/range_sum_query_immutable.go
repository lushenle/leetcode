package leetcode

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return NumArray{prefixSum: nums}
}

func (na *NumArray) SumRange(left int, right int) int {
	if left > 0 {
		return na.prefixSum[right] - na.prefixSum[left-1]
	}

	return na.prefixSum[right]
}
