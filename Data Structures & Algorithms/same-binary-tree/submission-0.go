/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	stackP := []*TreeNode{p}
	stackQ := []*TreeNode{q}
	for len(stackP) > 0 && len(stackQ) > 0 {
		curP := stackP[len(stackP)-1]
		curQ := stackQ[len(stackQ)-1]
		stackP = stackP[:len(stackP)-1]
		stackQ = stackQ[:len(stackQ)-1]
		if curP == nil && curQ == nil {
			continue
		}
		if curP == nil || curQ == nil || curP.Val != curQ.Val{
			return false
		}
		stackP = append(stackP, curP.Right, curP.Left)
		stackQ = append(stackQ, curQ.Right, curQ.Left)
	}
	if len(stackP) > 0 || len(stackQ) > 0 {
		return false
	}
	return true
}
