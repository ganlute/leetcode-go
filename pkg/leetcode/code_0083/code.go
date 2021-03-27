package code_0083

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pHead := head
	for pHead.Next != nil {
		if pHead.Val == pHead.Next.Val {
			pHead.Next = pHead.Next.Next
		} else {
			pHead = pHead.Next
		}
	}
	return head
}
