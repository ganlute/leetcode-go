package code_0234

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	p := head
	nodeMap := make(map[int]*ListNode, 0)
	i := 0
	for p != nil {
		nodeMap[i] = p
		i++
		p = p.Next
	}
	for j := 0; j <= i/2; j++ {
		if nodeMap[j].Val != nodeMap[i-1-j].Val {
			return false
		}
	}
	return true
}
