/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	head2 := slow.Next
	var prev *ListNode
	slow.Next = nil
	for head2 != nil {
		head2.Next, head2, prev = prev, head2.Next, head2
	}
	head2 = prev

	for head2 != nil {
		head, head2, head.Next, head2.Next = head.Next, head2.Next, head2, head.Next
	}
}
