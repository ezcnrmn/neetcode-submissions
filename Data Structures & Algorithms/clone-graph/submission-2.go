/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	queue := []*Node{node}
	seen := make(map[*Node]*Node)
	seen[node] = &Node{Val: node.Val}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		clone := seen[cur]

		for _, n := range cur.Neighbors {
			cloneN, ok := seen[n]
			if !ok {
				cloneN = &Node{Val: n.Val}
				seen[n] = cloneN
				queue = append(queue, n)
			}
			clone.Neighbors = append(clone.Neighbors, cloneN)
		}
	}
	return seen[node]
}

