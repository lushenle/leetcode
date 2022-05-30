package leetcode

import "github.com/lushenle/leetcode/structures"

type TreeNode = structures.TreeNode

func sumRootToLeaf(root *TreeNode) int {
	var dfs func(node *TreeNode, cur int) int
	dfs = func(node *TreeNode, cur int) int {
		if node == nil {
			return 0
		}

		next := (cur << 1) + node.Val
		if node.Left == nil && node.Right == nil {
			return next
		}

		return dfs(node.Left, next) + dfs(node.Right, next)
	}

	return dfs(root, 0)
}
