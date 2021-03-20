package code_19

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 1
	lenNode := head
	for lenNode.Next != nil {
		lenNode = lenNode.Next
		length = length + 1
	}
	n = length - n + 1 // 顺着第n个

	tHead := &ListNode{
		Next: head,
	}
	pHead := tHead
	for n > 1 {
		pHead = pHead.Next
		n = n - 1
	}
	pHead.Next = pHead.Next.Next
	return tHead.Next
}