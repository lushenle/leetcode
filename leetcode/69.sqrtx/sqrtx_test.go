package leetcode

import (
	"fmt"
	"testing"
)

type parameter struct {
	num int
}

type result struct {
	res int
}

type question struct {
	parameter
	result
}

func Test_mySqrt(t *testing.T) {
	qs := []question{
		{
			parameter{8},
			result{2},
		},
		{
			parameter{9},
			result{3},
		},
		{
			parameter{6},
			result{2},
		},
	}

	for _, q := range qs {
		r, p := q.result, q.parameter
		fmt.Printf("Input: %v, Output: %v, Expected: %v\n", p, mySqrt(p.num), r.res)
	}
}
