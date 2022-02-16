package leetcode

func checkWays(pairs [][]int) int {
	adj := map[int]map[int]bool{}
	for _, pair := range pairs {
		if adj[pair[0]] == nil {
			adj[pair[0]] = map[int]bool{}
		}
		if adj[pair[1]] == nil {
			adj[pair[1]] = map[int]bool{}
		}
		adj[pair[0]][pair[1]] = true
		adj[pair[1]][pair[0]] = true
	}
	n := len(adj)
	root := 0
	for node, neighbors := range adj {
		if len(neighbors) == n-1 {
			root = node
			break
		}
	}
	if root == 0 {
		return 0
	}
	res := 1
	for node, neighbors := range adj {
		if node == root {
			continue
		}
		nodeDegree := len(neighbors)
		nodeParent := 0
		parentDegree := n
		for neighbor := range neighbors {
			if len(adj[neighbor]) >= nodeDegree && len(adj[neighbor]) < parentDegree {
				nodeParent = neighbor
				parentDegree = len(adj[nodeParent])
			}
		}
		if nodeParent == 0 {
			return 0
		}

		for neighbor := range neighbors {
			if neighbor == nodeParent {
				continue
			}
			if !adj[nodeParent][neighbor] {
				return 0
			}
		}
		if nodeDegree == parentDegree {
			res = 2
		}

	}
	return res
}
