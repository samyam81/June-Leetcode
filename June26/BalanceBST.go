package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	values := make([]int, 0)
	dfs(root, &values)

	return buildBalancedBST(values, 0, len(values)-1)
}


func dfs(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}

	dfs(node.Left, values)
	*values = append(*values, node.Val)
	dfs(node.Right, values)
}

func buildBalancedBST(values []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	root := &TreeNode{
		Val:   values[mid],
		Left:  nil,
		Right: nil,
	}

	root.Left = buildBalancedBST(values, start, mid-1)
	root.Right = buildBalancedBST(values, mid+1, end)

	return root
}
