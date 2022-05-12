package leetcode

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	var cc Codec
	treeNode := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	fmt.Println(cc.serialize(treeNode))
	fmt.Println(cc.deserialize("213"))
}
