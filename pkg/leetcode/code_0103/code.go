package code_0103

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	return zigzagLevelOrderWithFlag(root, false)
}

// flag = true, 从左往右
// flag = false, 从右往左
func zigzagLevelOrderWithFlag(root *TreeNode, flag bool) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	result = append(result, []int{root.Val})
	left := zigzagLevelOrderWithFlag(root.Left, !flag)
	right := zigzagLevelOrderWithFlag(root.Right, !flag)

	for i := 0; i < len(left) || i < len(right); i++ {
		tResult := make([]int, 0)
		if flag {
			if i < len(left) {
				tResult = append(tResult, left[i]...)
			}
			if i < len(right) {
				tResult = append(tResult, right[i]...)
			}
		} else {
			if i < len(right) {
				tResult = append(tResult, right[i]...)
			}
			if i < len(left) {
				tResult = append(tResult, left[i]...)
			}
		}
		result = append(result, tResult)
		flag = !flag
	}
	return result
}
