package code_141


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	pointMap := make(map[*ListNode]uint8, 0 )

	for head != nil {
		_, ok := pointMap[head]
		if ok {
			return true
		}
		pointMap[head] = 0
		head = head.Next
	}
	return false
}