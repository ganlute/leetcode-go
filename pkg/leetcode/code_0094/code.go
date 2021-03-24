package code_0094

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 中序遍历首先遍历左子树，然后访问根结点，最后遍历右子树。
// 后序遍历首先遍历左子树，然后遍历右子树，最后访问根结点。
// 前序遍历首先访问根结点，然后遍历左子树，最后遍历右子树。
func inorderTraversal(root *TreeNode) []int {
	list := make([]int,0)
	if root == nil {
		return list
	}
	list = append(list, inorderTraversal(root.Left)...)
	list = append(list, root.Val)
	list = append(list, inorderTraversal(root.Right)...)
	return list
}