package leetcode

import "math/rand"

type RandomizedSet struct {
	nums  []int
	index map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{index: map[int]int{}, nums: []int{}}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.index[val]; ok {
		return false
	}

	rs.nums = append(rs.nums, val)
	rs.index[val] = len(rs.nums) - 1
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	if id, ok := rs.index[val]; ok {
		last := len(rs.nums) - 1
		rs.nums[id] = rs.nums[last]
		rs.index[rs.nums[id]] = id
		rs.nums = rs.nums[:last]
		delete(rs.index, val)
		return true
	}

	return false
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}
