package leetcode

import "testing"

func TestDetectSquares(t *testing.T) {
	ds := Constructor()
	ds.Add([]int{3, 10})
	ds.Add([]int{11, 2})
	ds.Add([]int{3, 2})
	ds.Add([]int{11, 10})

	n := ds.Count([]int{11, 10})
	if n != 1 {
		t.Errorf("count should be 1, got %d", n)
	}

	n2 := ds.Count([]int{14, 8})
	if n2 != 0 {
		t.Errorf("count should be 0, got %d", n)
	}

	ds.Add([]int{11, 2})

	n3 := ds.Count([]int{11, 10})
	if n3 != 2 {
		t.Errorf("count should be 2, got %d", n)
	}
}
