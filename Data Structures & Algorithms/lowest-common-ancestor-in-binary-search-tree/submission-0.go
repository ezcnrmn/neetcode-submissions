/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Val == p.Val || cur.Val == q.Val {
			return cur
		}
		if (cur.Val > p.Val && cur.Val < q.Val) || (cur.Val < p.Val && cur.Val > q.Val) {
			return cur
		}
		if cur.Val > p.Val && cur.Val > q.Val {
			stack = append(stack, cur.Left)
		} else {
			stack = append(stack, cur.Right)
		}
	}
	return nil
}
