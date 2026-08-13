/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type StackItem struct {
	Node        *TreeNode
	Parent      *StackItem
	ChildLength []int
}

func NewStackItem(node *TreeNode, parent *StackItem) *StackItem {
	return &StackItem{Node: node, Parent: parent, ChildLength: make([]int, 0, 2)}
}

func isBalanced(root *TreeNode) bool {
	dummy := &StackItem{}
	stack := []*StackItem{NewStackItem(root, dummy)}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		if cur.Node == nil {
			stack = stack[:len(stack)-1]
			cur.Parent.ChildLength = append(cur.Parent.ChildLength, 0)
			continue
		}

		if len(cur.ChildLength) == 2 {
			stack = stack[:len(stack)-1]
			maxLength := max(cur.ChildLength[0], cur.ChildLength[1]) 
			minLength := min(cur.ChildLength[0], cur.ChildLength[1]) 
			if maxLength - minLength > 1 { 
				return false
			}
			cur.Parent.ChildLength = append(cur.Parent.ChildLength, maxLength+1)
			continue
		}

		stack = append(stack, NewStackItem(cur.Node.Right, cur))
		stack = append(stack, NewStackItem(cur.Node.Left, cur))
	}
	return true
}
