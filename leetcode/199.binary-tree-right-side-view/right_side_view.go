package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	// 初始化
	var result []int
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) != 0 {
		// 当前层的节点数
		size := len(queue)
		// 当前层的节点
		var nodes []*TreeNode
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			nodes = append(nodes, node)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 当前层的最右节点
		result = append(result, nodes[len(nodes)-1].Val)
	}

	return result
}
