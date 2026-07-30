/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l, r := length(root.Left), length(root.Right)
	return math.Abs(l - r) <= 1
}

func length(n *TreeNode) float64 {
	if n == nil {
		return 0
	}
	return 1 + max(length(n.Left), length(n.Right))
}