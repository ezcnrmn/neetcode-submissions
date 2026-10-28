/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}

	var thisCompare bool
	if root.Val == subRoot.Val {
		thisCompare = isSame(root, subRoot)
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot) || thisCompare
}

func isSame(r1, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil || r1.Val != r2.Val {
		return false
	}

	return isSame(r1.Left, r2.Left) && isSame(r1.Right, r2.Right)
}