package code_0023

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else {
		return merge(lists[0], mergeKLists(lists[1:]))
	}
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{
		Next: nil,
	}
	p := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			p = p.Next
			list1 = list1.Next
		} else {
			p.Next = list2
			p = p.Next
			list2 = list2.Next
		}
	}
	if list1 == nil {
		p.Next = list2
	} else {
		p.Next = list1
	}
	return head.Next
}
