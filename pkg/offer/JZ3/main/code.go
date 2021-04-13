package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param head ListNode类
 * @return int整型一维数组
 */
func printListFromTailToHead(head *ListNode) []int {
	// write code here
	tmpArr := make([]int, 0)
	for head != nil {
		tmpArr = append(tmpArr, head.Val)
		head = head.Next
	}
	return reverse(tmpArr)
}
func reverse(in []int) []int {
	l := len(in)
	for i := 0; i < l/2; i++ {
		in[i], in[l-1-i] = in[l-1-i], in[i]
	}
	return in
}
