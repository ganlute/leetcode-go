package code_0102

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	result = append(result, []int{root.Val})

	left := levelOrder(root.Left)
	right := levelOrder(root.Right)

	for i := 0; i < len(left) || i < len(right); i++ {
		t := make([]int, 0)
		if i < len(left) {
			t = append(t, left[i]...)
		}
		if i < len(right) {
			t = append(t, right[i]...)
		}
		result = append(result, t)

	}
	return result
}
