package main

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// Traverse right subtree first
		dfs(node.Right)
		// Update sum and node value
		sum += node.Val
		node.Val = sum
		// Traverse left subtree
		dfs(node.Left)
	}

	dfs(root)
	return root
}