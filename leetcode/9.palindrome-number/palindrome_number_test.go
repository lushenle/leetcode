package leetcode

import (
	"fmt"
	"testing"
)

type parameter struct {
	num int
}

type result struct {
	res bool
}

type question struct {
	parameter
	result
}

func Test_mySqrt(t *testing.T) {
	qs := []question{
		{
			parameter{121},
			result{true},
		},
		{
			parameter{12321},
			result{true},
		},
		{
			parameter{-123},
			result{false},
		},
	}

	for _, q := range qs {
		r, p := q.result, q.parameter
		fmt.Printf("Input: %v, Output: %v, Expected: %v\n", p.num, isPalindrome(p.num), r.res)
	}
}
