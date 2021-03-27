package code_0082

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	tHead := &ListNode{
		Next: head,
	}
	pHead := tHead

	for pHead != nil {
		if pHead.Next != nil && pHead.Next.Next != nil && pHead.Next.Val == pHead.Next.Next.Val {
			// start to delete
			qHead := pHead.Next
			tValue := qHead.Val
			for qHead != nil {
				if qHead.Val == tValue {
					qHead = qHead.Next
				} else {
					break
				}
			}
			pHead.Next = qHead
		} else {
			pHead = pHead.Next
		}
	}
	return tHead.Next
}
