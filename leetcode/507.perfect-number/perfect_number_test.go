package leetcode

import (
	"fmt"
	"log"
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
			result{false},
		},
	}

	for _, q := range qs {
		r, p := q.result, q.parameter
		fmt.Printf("Input: %v, Output: %v, Expected: %v\n", p.num, checkPerfectNumber(p.num), r.res)
		if checkPerfectNumber(p.num) != r.res {
			log.Fatalf("testing failed: expected: %v", r.res)
		}
	}
}
