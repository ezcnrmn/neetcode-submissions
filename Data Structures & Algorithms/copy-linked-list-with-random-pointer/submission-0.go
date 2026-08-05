/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	mem := make(map[*Node]*Node)
	dummy := &Node{}
	cur1, cur2 := head, dummy

	for cur1 != nil {
		node := &Node{Val: cur1.Val}
		mem[cur1] = node
		cur2.Next = node
		cur2 = cur2.Next
		cur1 = cur1.Next
	}

	cur1, cur2 = head, dummy.Next
	for cur1 != nil {
		cur2.Random = mem[cur1.Random]
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	return dummy.Next
}
