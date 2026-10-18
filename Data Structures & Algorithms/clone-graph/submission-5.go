/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	
	clones := make(map[*Node]*Node)
	visited := make(map[*Node]struct{})

	stack := []*Node{node}

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if _, ok := clones[cur]; !ok {
			clones[cur] = &Node{Val: cur.Val}
		}
		clone := clones[cur]
		visited[cur] = struct{}{}

		for _, nbr := range cur.Neighbors {
			if _, ok := clones[nbr]; !ok {
				clones[nbr] = &Node{Val: nbr.Val}
				stack = append(stack, nbr)
			}
			nbrClone := clones[nbr]
			clone.Neighbors = append(clone.Neighbors, nbrClone)
		}
	}

	return clones[node]
}
