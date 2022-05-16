package leetcode

import "github.com/lushenle/leetcode/structures"

type TreeNode = structures.TreeNode

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > p.Val {
		if ans := inorderSuccessor(root.Left, p); ans != nil {
			return ans
		}
		return root
	}

	return inorderSuccessor(root.Right, p)
}
