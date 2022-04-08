package leetcode

// Reference: 429

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) (ans []int) {
	nodes := []*Node{root, nil}

	for l := len(nodes); l > 0; l = len(nodes) {
		node := nodes[l-1]
		nodes = nodes[:l-1]

		if node != nil {
			ans = append(ans, node.Val)
			for i := len(node.Children) - 1; i >= 0; i-- {
				nodes = append(nodes, node.Children[i])
			}
		}
	}

	return
}
