package code_24

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// H @ @ @ @ @
// p
// t   p
// H --> l1 --> l2 --> l3
// H --> l2 --> l1 -->l3
func swapPairs(head *ListNode) *ListNode {
	tHead := &ListNode{
		Next: head,
	}
	pNode := tHead // 始终指向当前处理逻辑的前一个节点
	for pNode.Next != nil && pNode.Next.Next != nil { // 符合处理条件
		// L0 L1 L2 L3
		L0 := pNode
		L1 := pNode.Next
		L2 := pNode.Next.Next
		L3 := pNode.Next.Next.Next
		// L0 L2 L1 L3
		L0.Next = L2
		L2.Next = L1
		L1.Next = L3
		pNode = L1
	}
	return tHead.Next
}