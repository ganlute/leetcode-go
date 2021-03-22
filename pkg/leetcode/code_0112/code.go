package code_0112

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return dfs(root, targetSum)
}

func dfs(root *TreeNode, targetSum int) bool {
	if root == nil {
		if targetSum == 0 {
			return true
		} else {
			return false
		}
	}
	if root.Right != nil && root.Left != nil {
		return dfs(root.Right, targetSum-root.Val) || dfs(root.Left, targetSum-root.Val)
	}
	if root.Right != nil && root.Left == nil {
		return dfs(root.Right, targetSum-root.Val)
	}
	if root.Right == nil && root.Left != nil {
		return dfs(root.Left, targetSum-root.Val)
	}
	return dfs(nil, targetSum-root.Val)
}
