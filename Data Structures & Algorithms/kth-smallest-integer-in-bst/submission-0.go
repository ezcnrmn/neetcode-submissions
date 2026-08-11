/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	var f func(*TreeNode)
	var result int
	var counter int
	f = func(node *TreeNode) {
		if node == nil {
			return 
		}
		f(node.Left)
		counter++
		if counter == k {
			result = node.Val
		}
		f(node.Right)
	}
	f(root)
	return result
}
