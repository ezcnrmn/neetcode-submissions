/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	switch {
	case root.Val == p.Val || root.Val == q.Val,
		(root.Val > p.Val && root.Val < q.Val) || (root.Val < p.Val && root.Val > q.Val):  
		return root
	case root.Val > p.Val && root.Val > q.Val:
		return lowestCommonAncestor(root.Left, p, q)
	default:
		return lowestCommonAncestor(root.Right, p, q)
	}
	return nil
}