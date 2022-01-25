package leetcode

import (
	"fmt"
	"testing"
)

type parameter struct {
	nums   []int
	target int
}

type result struct {
	res []int
}

type question struct {
	parameter
	result
}

func Test_twoSum(t *testing.T) {
	qs := []question{
		{
			parameter{[]int{1, 2, 3, 4, 5}, 8},
			result{[]int{2, 4}},
		},
		{
			parameter{[]int{2, 7, 11, 15}, 9},
			result{[]int{0, 1}},
		},
	}

	for _, q := range qs {
		r, p := q.result, q.parameter
		fmt.Printf("Input: %v, Output: %v, Expected: %v\n", p, twoSum(p.nums, p.target), r.res)
	}
}
