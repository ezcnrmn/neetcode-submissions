/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	var prev *ListNode
	for cur.Next != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	cur.Next = prev
	return cur
}
