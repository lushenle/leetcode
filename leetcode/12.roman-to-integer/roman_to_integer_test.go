package leetcode

import (
	"fmt"
	"testing"
)

type parameter struct {
	roman string
}

type result struct {
	res int
}

type question struct {
	parameter
	result
}

func Test_romanToInt(t *testing.T) {
	qs := []question{
		{
			parameter{"LVIII"},
			result{58},
		},
		{
			parameter{"III"},
			result{3},
		},
		{
			parameter{"MCMXICIVI"},
			result{2014},
		},
	}

	for _, q := range qs {
		r, p := q.result, q.parameter
		fmt.Printf("Input: %v, Output: %v, Expected: %v\n", p.roman, romanToInt(p.roman), r.res)
	}
}
