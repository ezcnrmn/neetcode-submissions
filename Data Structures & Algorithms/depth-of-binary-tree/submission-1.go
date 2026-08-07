/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type StackItem struct {
	Node  *TreeNode
	Depth int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var stack []StackItem
	var result int
	stack = append(stack, StackItem{Node: root, Depth: 1})
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur, depth := item.Node, item.Depth
		if cur != nil {
			result = max(result, depth)
			stack = append(stack, StackItem{Node: cur.Right, Depth: depth+1})
			stack = append(stack, StackItem{Node: cur.Left, Depth: depth+1})
		}
	}
	return result
}
