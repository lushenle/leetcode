package leetcode

import "github.com/lushenle/leetcode/structures"

type TreeNode = structures.TreeNode

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil && root.Left.Val != root.Val {
		return false
	}

	if root.Right != nil && root.Right.Val != root.Val {
		return false
	}

	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}

func isUnivalTree1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	val := root.Val
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node != nil {
			return node.Val == val && dfs(node.Left) && dfs(node.Right)
		}
		return true
	}

	return dfs(root)
}
