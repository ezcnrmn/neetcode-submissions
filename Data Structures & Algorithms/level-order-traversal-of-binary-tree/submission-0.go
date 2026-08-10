/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		level := make([]int, 0, length)
		for range length {
			cur := queue[0]
			queue = queue[1:]
			if cur == nil {
				continue
			}
			level = append(level, cur.Val)
			queue = append(queue, cur.Left, cur.Right)
		}
		if len(level) != 0 {
			result = append(result, level)
		}
	}

	return result
}
