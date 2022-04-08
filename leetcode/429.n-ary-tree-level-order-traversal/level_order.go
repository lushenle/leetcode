package leetcode

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var (
		ans  [][]int
		temp []int
	)

	if root == nil {
		return ans
	}

	queue := []*Node{root, nil}
	for len(queue) > 1 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			queue = append(queue, nil)
			ans = append(ans, temp)
			temp = []int{}
		} else {
			temp = append(temp, node.Val)
			if len(node.Children) > 0 {
				queue = append(queue, node.Children...)
			}
		}
	}

	ans = append(ans, temp)
	return ans
}
