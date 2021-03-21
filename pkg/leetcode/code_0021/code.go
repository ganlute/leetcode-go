package code_0021

// Merge two sorted linked lists and return it as a sorted list.
// The list should be made by splicing together the nodes of the first two lists.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	pre := l
	for l1 != nil || l2 != nil {
		if l1 == nil {
			pre.Next = l2
			l2 = l2.Next
			pre = pre.Next
		}else if l2 == nil {
			pre.Next = l1
			l1 = l1.Next
			pre = pre.Next
		}else if l2.Val < l1.Val{
			pre.Next = l2
			l2 = l2.Next
			pre = pre.Next
		}else {
			pre.Next = l1
			l1 = l1.Next
			pre = pre.Next
		}
	}
	return l.Next
}
