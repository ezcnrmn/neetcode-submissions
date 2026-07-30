/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	cur := head
	mem := make(map[*ListNode]bool)
	for cur != nil {
		if mem[cur] {
			return true
		}
		mem[cur] = true
		cur = cur.Next
	}
	return false
}
