package leetcode

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	ranges := [][]int{
		{0, 2},
		{2, 5},
		{0, 5},
	}

	e := []int{1, -1, -3}

	for i, s := range ranges {
		if res := obj.SumRange(s[0], s[1]); res != e[i] {
			t.Errorf("Expected: %d, Output: %d", e[i], res)
		}
	}
}
