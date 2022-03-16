package leetcode

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AllOne(t *testing.T) {
	ast := assert.New(t)
	a := Constructor()

	ast.Equal("", a.GetMaxKey())
	ast.Equal("", a.GetMinKey())

	a.Inc("666")
	ast.Equal("666", a.GetMinKey())
	ast.Equal("666", a.GetMaxKey())

	a.Dec("666")
	ast.Equal("", a.GetMaxKey())
	ast.Equal("", a.GetMinKey())

	maxKeys := []string{"3", "4", "5"}
	minKeys := []string{"0", "1", "2"}

	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			a.Inc(strconv.Itoa(j))
		}
	}

	for _, key := range maxKeys {
		a.Inc(key)
		a.Inc(key)
	}

	for _, key := range minKeys {
		a.Dec(key)
		a.Dec(key)
	}

	ast.Contains(maxKeys, a.GetMaxKey(), "get max key error")
	ast.Contains(minKeys, a.GetMinKey(), "get min key error")

	a.Dec("123")
}
