/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	cur := slow.Next
	var prev *ListNode
	for cur != nil{
		cur.Next, prev, cur = prev, cur, cur.Next
	}

	p1, p2 := head, prev
	for p2 != nil {
		temp1, temp2 := p1.Next, p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p1, p2 = temp1, temp2
	}
	if p1 != nil {
		p1.Next = nil
	}

	return 
}
