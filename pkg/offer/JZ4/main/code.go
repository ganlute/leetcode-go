package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	// write code here
	if len(pre) == 0 {
		return nil
	}
	i := 0
	for {
		if vin[i] == pre[0] {
			break
		} else {
			i++
		}
	}
	treeNode := &TreeNode{
		Val:   pre[0],
		Left:  reConstructBinaryTree(pre[1:i+1], vin[0:i]),
		Right: reConstructBinaryTree(pre[i+1:], vin[i+1:]),
	}
	return treeNode
}
