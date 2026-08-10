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
	Max  int
}

func goodNodes(root *TreeNode) int {
	stack := []StackItem{{Node: root, Max: 0}}
	var result int
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Node == nil {
			continue
		}
		if cur.Node.Val >= cur.Max {
			result++
		}
		curMax := max(cur.Max, cur.Node.Val)
		stack = append(stack, StackItem{Node: cur.Node.Right, Max: curMax})
		stack = append(stack, StackItem{Node: cur.Node.Left, Max: curMax})
	}
	return result
}
