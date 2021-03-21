package code_0002

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersWithInt(l1, l2, 0)
}

func addTwoNumbersWithInt(l1 *ListNode, l2 *ListNode, value int) *ListNode {
	if l1 == nil && l2 == nil && value == 0{
		return nil
	}

	v1 := 0
	var n1  *ListNode

	v2 := 0
	var n2 *ListNode

	if l1 != nil {
		v1 = l1.Val
		n1 = l1.Next
	}
	if l2 != nil {
		v2 = l2.Val
		n2 = l2.Next
	}
	value = value + v1 + v2


	listNode := &ListNode{
		Val:  value % 10,
		Next: addTwoNumbersWithInt(n1, n2, value / 10),
	}

	return listNode
}