package leetcode

import (
	"fmt"
	"testing"
)

type result struct {
	res bool
}

type parameter struct {
	num int
}

type question struct {
	parameter
	result
}

func Test_perfectNumber(t *testing.T) {
	qs := []question{
		{
			parameter{6},
			result{true},
		},
		{
			parameter{28},
			result{true},
		},
		{
			parameter{496},
			result{true},
		},
		{
			parameter{8128},
			result{true},
		},
		{
			parameter{33550336},
			result{true},
		},
		{
			parameter{27},
			result{true},
		},
	}

	for _, q := range qs {
		_, p := q.result, q.parameter
		fmt.Printf("Input: %v, Output: %v\n", p, checkPerfectNumber(p.num))
	}
}
