package code_0101

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return issymmetric(root.Left, root.Right)
}

func issymmetric(tNode1, tNode2 *TreeNode) bool {
	if tNode1 == nil && tNode2 == nil {
		return true
	}
	if tNode1 == nil && tNode2 != nil {
		return false
	}
	if tNode1 != nil && tNode2 == nil {
		return false
	}
	return (tNode1.Val == tNode2.Val) &&
		issymmetric(tNode1.Left, tNode2.Right) &&
		issymmetric(tNode1.Right, tNode2.Left)
}
