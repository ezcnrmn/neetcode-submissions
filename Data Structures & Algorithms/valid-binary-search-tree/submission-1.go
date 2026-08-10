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
	Min  int
	Max  int
}

func isValidBST(root *TreeNode) bool {
	stack := []StackItem{{Node: root, Min: -1000000001, Max: 1000000001}}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Node == nil {
			continue
		}
		if cur.Node.Val <= cur.Min || cur.Node.Val >= cur.Max {
			return false
		}
		stack = append(stack, StackItem{Node: cur.Node.Right, Min: cur.Node.Val, Max: cur.Max})
		stack = append(stack, StackItem{Node: cur.Node.Left, Min: cur.Min, Max: cur.Node.Val})
	}
	return true
}
