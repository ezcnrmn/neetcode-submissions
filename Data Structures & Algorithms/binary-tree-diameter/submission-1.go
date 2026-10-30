/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	var result int
	var f func(*TreeNode) int
	f = func (node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := f(node.Left)
		right := f(node.Right)
		result = max(result, left + right)
		return max(left, right) + 1
	}
	f(root)
	return result
}
