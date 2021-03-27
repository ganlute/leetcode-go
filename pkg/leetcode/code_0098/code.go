package code_0098

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 中序遍为升序
func isValidBST(root *TreeNode) bool {
	result := scanTreeNode(root)
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}

func scanTreeNode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	result = append(result, scanTreeNode(root.Left)...)
	result = append(result, root.Val)
	result = append(result, scanTreeNode(root.Right)...)
	return result
}
