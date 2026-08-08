/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	var f func(*TreeNode) int
	result := true
	f = func (node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := f(node.Left)
		right := f(node.Right)
		left, right = max(left, right), min(left, right)
		if left - right > 1 {
			result = false
		}
		return left + 1
	}
	f(root)
	return result
}

