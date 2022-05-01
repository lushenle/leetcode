package leetcode

import "github.com/lushenle/leetcode/structures"

type TreeNode = structures.TreeNode

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	nums1 := inorderTraversal(root1)
	nums2 := inorderTraversal(root2)
	nums1 = append(nums1, make([]int, len(nums2))...)
	merge(nums1, len(nums1)-len(nums2), nums2, len(nums2))
	return nums1
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	inorder(root, &result)
	return result
}

func inorder(root *TreeNode, output *[]int) {
	if root != nil {
		inorder(root.Left, output)
		*output = append(*output, root.Val)
		inorder(root.Right, output)
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	i := m - 1
	j := n - 1
	k := m + n - 1

	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}

	for ; j >= 0; k-- {
		nums1[k] = nums2[j]
		j--
	}
}
