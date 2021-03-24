package code_0092

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 简单易懂做题
// 1、先找到pLeft的位置，固定
// 2、找到pRight的位置，固定
// 3、把pLeft 和 pRight之间的node保存
// 4、反转
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	tHead := &ListNode{
		Next: head,
	}
	mapNode := make(map[int]*ListNode, right-left+1)
	pLeft := tHead
	// 先找到pLeft的位置，固定
	for i := 1; i < left; i++ {
		pLeft = pLeft.Next
	}
	pRight := pLeft.Next
	for i := left; i <= right; i++ {
		mapNode[i-left] = pRight
		pRight = pRight.Next
	}
	for i := 0; i < right-left; i++ {
		mapNode[i+1].Next = mapNode[i]
	}
	pLeft.Next = mapNode[right-left]
	mapNode[0].Next = pRight
	return tHead.Next
}
