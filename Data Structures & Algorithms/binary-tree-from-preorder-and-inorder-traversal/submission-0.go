/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import (
	"slices"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	node := &TreeNode{Val: preorder[0]}
	i := slices.Index(inorder, node.Val)
	left, right := inorder[:i], inorder[i+1:]
	node.Left = buildTree(preorder[1:], left)
	node.Right = buildTree(preorder[len(left)+1:], right)
	return node
}
