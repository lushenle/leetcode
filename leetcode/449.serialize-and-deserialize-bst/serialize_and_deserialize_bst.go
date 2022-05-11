package leetcode

import "github.com/lushenle/leetcode/structures"

type TreeNode = structures.TreeNode

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

func insert(node *TreeNode, value int) {
	if value <= node.Val {
		if node.Left == nil {
			node.Left = &TreeNode{Val: value}
		} else {
			insert(node.Left, value)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{Val: value}
		} else {
			insert(node.Right, value)
		}
	}
}

func (cc *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var nodes []rune
	var preOrder func(*TreeNode)

	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		nodes = append(nodes, rune(node.Val))

		preOrder(node.Left)
		preOrder(node.Right)
	}

	preOrder(root)

	return string(nodes)
}

func (cc *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	root := &TreeNode{}

	for i, n := range data {
		if i == 0 {
			root.Val = int(n)
		} else {
			insert(root, int(n))
		}
	}

	return root
}
