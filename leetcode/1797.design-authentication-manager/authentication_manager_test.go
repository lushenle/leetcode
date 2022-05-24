package leetcode

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	am := Constructor(5)
	am.timeToLive = 5
	am.Renew("aaa", 1)
	am.Generate("aaa", 2)
	am.CountUnexpiredTokens(6)
	am.Generate("bbb", 7)
	am.Renew("aaa", 8)
	am.Renew("bbb", 10)
	am.CountUnexpiredTokens(15)
	fmt.Println(am)
}
