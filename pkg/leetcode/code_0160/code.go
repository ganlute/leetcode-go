package code_0160

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 遍历两个链表，先把长度对其，超出部分截断
	delta := length(headA) - length(headB)
	if delta > 0 {
		// 截断A
		for delta > 0 {
			delta = delta - 1
			headA = headA.Next
		}
	} else if delta < 0 {
		// 截断B
		for delta < 0 {
			delta = delta + 1
			headB = headB.Next
		}
	}
	// 寻找公共部分
	pA := headA
	pB := headB
	for pA != nil {
		if pA == pB {
			return pA
		}
		pA = pA.Next
		pB = pB.Next
	}
	return nil
}
func length(head *ListNode) int {
	pHead := head
	count := 0
	for pHead != nil {
		count++
		pHead = pHead.Next
	}
	return count
}
