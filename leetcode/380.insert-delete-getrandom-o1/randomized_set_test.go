package leetcode

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	rs := Constructor()

	if ok := rs.Insert(1); !ok {
		t.Errorf("should be insert 1, but not")
	}

	if ok := rs.Insert(1); ok {
		t.Errorf("should be not insert 1")
	}

	if ok := rs.Remove(3); ok {
		t.Errorf("val 3 is not exist,catnot remove")
	}

	rs.Insert(2)
	rs.Insert(3)
	rs.Insert(4)

	if ok := rs.Remove(3); !ok {
		t.Errorf("should be remove val 3")
	}

	fmt.Println(rs.GetRandom())

}
