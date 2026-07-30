/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var head, result *ListNode
	if list1.Val < list2.Val {
		head, result = list1, list1
		list1 = list1.Next
	} else {
		head, result = list2, list2
		list2 = list2.Next
	}
	
	for list1 != nil || list2 != nil {
		if list1 == nil || (list2 != nil && list2.Val < list1.Val) {
			result.Next = list2
			list2 = list2.Next
		} else {
			result.Next = list1
			list1 = list1.Next
		}
		result = result.Next
	}

	return head
}
