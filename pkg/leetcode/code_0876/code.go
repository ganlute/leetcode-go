package code_0876

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	// 第一遍扫描确定链表长度
	pHead := head
	length := 0
	for pHead != nil {
		pHead = pHead.Next
		length++
	}
	// 第二遍扫描拿到中间节点
	tHead := &ListNode{
		Next: head,
	}
	pHead = tHead

	for i := 1; i <= length/2+1; i++ {
		pHead = pHead.Next
	}
	return pHead
}
