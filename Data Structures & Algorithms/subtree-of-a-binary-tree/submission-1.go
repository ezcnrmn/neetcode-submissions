/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur == nil {
			continue
		}
		if cur.Val == subRoot.Val && isSame(cur, subRoot) {
			return true
		}
		stack = append(stack, cur.Left, cur.Right)
	}
	return false
}

func isSame(root1, root2 *TreeNode) bool {
	stack1 := []*TreeNode{root1}
	stack2 := []*TreeNode{root2}
	for len(stack1) > 0 && len(stack2) > 0 {
		cur1 := stack1[len(stack1)-1]
		cur2 := stack2[len(stack2)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = stack2[:len(stack2)-1]
		if cur1 == nil && cur2 == nil {
			continue
		}
		if cur1 != nil && cur2 != nil && cur1.Val == cur2.Val {
			stack1 = append(stack1, cur1.Left, cur1.Right)
			stack2 = append(stack2, cur2.Left, cur2.Right)
			continue
		}
		return false
	}
	return len(stack1) == 0 && len(stack2) == 0
}