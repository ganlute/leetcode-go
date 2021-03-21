package code_0206

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	pNode := head
	tNode := &ListNode{}
	for pNode != nil {
		qNode := pNode.Next
		pNode.Next = tNode.Next
		tNode.Next = pNode
		pNode = qNode
	}
	return tNode.Next
}