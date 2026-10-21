/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type StackItem struct {
	Node *TreeNode
	Seen bool
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []StackItem{{Node: root}}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cur.Node == nil {
			continue
		}

		if cur.Seen {
			k--
			if k == 0 {
				return cur.Node.Val
			}
		} else {
			stack = append(
				stack,
				StackItem{Node: cur.Node.Right},
				StackItem{Node: cur.Node, Seen: true},
				StackItem{Node: cur.Node.Left},
			)
		}
	}

	return -1
}
