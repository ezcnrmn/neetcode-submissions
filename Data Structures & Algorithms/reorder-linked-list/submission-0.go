/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	var mem []*ListNode
	cur := head
	for cur != nil {
		mem = append(mem, cur)
		cur = cur.Next
	}
	for i, j := 0, len(mem)-1; i < len(mem)/2; i, j = i+1, j-1 {
		mem[i].Next = mem[j]
		if i+1 != j {
			mem[j].Next = mem[i+1]
		}
	}
	mem[len(mem)/2].Next = nil
}
