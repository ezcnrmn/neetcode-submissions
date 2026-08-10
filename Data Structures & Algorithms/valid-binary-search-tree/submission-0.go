/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Left != nil {
			if cur.Val <= cur.Left.Val {
				return false
			}
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			if cur.Val >= cur.Right.Val {
				return false
			}
			stack = append(stack, cur.Right)
		}
	}
	return true
}
