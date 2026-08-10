/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	queue := []*TreeNode{root}
	var result []int
	var levelCounter int
	for len(queue) > 0 {
		length := len(queue)
		levelCounter++
		for range length {
			cur := queue[0]
			queue = queue[1:]
			if cur == nil {
				continue
			}
			if len(result) < levelCounter {
				result = append(result, cur.Val)
			}
			queue = append(queue, cur.Right, cur.Left)
		}
	}
	return result
}
