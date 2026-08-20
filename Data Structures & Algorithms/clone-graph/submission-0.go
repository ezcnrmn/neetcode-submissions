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

	stack := []*Node{node}
	seen := make(map[*Node]*Node)
	seen[node] = &Node{Val: node.Val}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		clone := seen[cur]

		for _, n := range cur.Neighbors {
			cloneN, ok := seen[n]
			if !ok {
				cloneN = &Node{Val: n.Val}
				seen[n] = cloneN
				stack = append(stack, n)
			}
			clone.Neighbors = append(clone.Neighbors, cloneN)
		}
	}
	return seen[node]
}

