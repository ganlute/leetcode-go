package code_0142

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]uint8, 0)
	pHead := head
	for pHead != nil {
		if _, ok := nodeMap[pHead]; !ok {
			nodeMap[pHead] = 1
		} else {
			return pHead
		}
		pHead = pHead.Next
	}
	return nil
}
