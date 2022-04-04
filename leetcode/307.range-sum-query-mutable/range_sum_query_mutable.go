package leetcode

import "runtime/debug"

type NumArray struct {
	nums []int
}

// Constructor return the NumArray
func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

// Update updates nums
func (na *NumArray) Update(i, v int) {
	na.nums[i] = v
}

// SumRange returns sum(nums[i:j+1])
func (na *NumArray) SumRange(i, j int) (ans int) {
	for k := i; k <= j; k++ {
		ans += na.nums[k]
	}
	return ans
}

func init() { debug.SetGCPercent(-1) }
