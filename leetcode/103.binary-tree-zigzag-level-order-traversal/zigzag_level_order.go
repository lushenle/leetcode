package leetcode

import (
	"github.com/lushenle/leetcode/structures"
)

// TreeNode definition for a binary tree node
type TreeNode = structures.TreeNode

// dfs
func zigzagLevelOrder(root *TreeNode) [][]int {
	var ans [][]int

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		// 新的 level
		if level >= len(ans) {
			ans = append(ans, []int{})
		}

		if level%2 == 0 {
			ans[level] = append(ans[level], root.Val)
		} else {
			temp := make([]int, len(ans[level])+1)
			temp[0] = root.Val
			copy(temp[1:], ans[level])
			ans[level] = temp
		}

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)

	return ans
}

// bfs
func zigzagLevelOrder1(root *TreeNode) [][]int {
	var ans [][]int

	if root == nil {
		return ans
	}

	q := []*TreeNode{root}
	var level []int
	var temp []*TreeNode
	size, i, j, flag := 0, 0, 0, false

	for len(q) > 0 {
		size = len(q)
		temp = []*TreeNode{}
		level = make([]int, size)
		j = size - 1

		for i = 0; i < size; i++ {
			root = q[0]
			q = q[1:]
			if !flag {
				level[i] = root.Val
			} else {
				level[j] = root.Val
				j--
			}

			if root.Left != nil {
				temp = append(temp, root.Left)
			}

			if root.Right != nil {
				temp = append(temp, root.Right)
			}
		}

		ans = append(ans, level)
		flag = !flag
		q = temp
	}

	return ans
}
