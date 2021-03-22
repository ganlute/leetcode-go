package code_0025

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	list := make([]*ListNode, k)
	i := 0
	for head != nil && i < k {
		list[i] = head
		head = head.Next
		i++
	}
	i--
	// 返回
	if i+1 != k {
		return list[0]
	} else {
		// 反转list
		for j := i; j > 0; j-- {
			list[j].Next = list[j-1]
		}
		list[0].Next = reverseKGroup(head, k)
		return list[i]
	}

}
