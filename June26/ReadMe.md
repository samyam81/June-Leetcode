# Balanced Binary Search Tree (BST) 

This Go package (`main`) provides functions to balance an unbalanced Binary Search Tree (BST) by converting it into a balanced BST.

## TreeNode Struct

The package defines a `TreeNode` struct representing a node in the BST:

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
```

## Function: balanceBST

The `balanceBST` function is the main entry point which balances the input BST. It takes the root node of the BST (`root *TreeNode`) and returns the root node of the balanced BST (`*TreeNode`).

### Explanation

The `balanceBST` function works in two main steps:

1. **Depth-First Search (DFS) Traversal (`dfs`)**:
   - This helper function recursively traverses the input BST (`node *TreeNode`).
   - It performs an in-order traversal (`Left`, `Root`, `Right`) and collects node values into the `values` slice.

2. **Building the Balanced BST (`buildBalancedBST`)**:
   - After collecting all node values in sorted order using DFS, the `buildBalancedBST` function is called.
   - It constructs a balanced BST recursively using a divide-and-conquer approach.
   - It calculates the middle index (`mid`) of the current range (`start` to `end`) of `values` and sets `values[mid]` as the value of the current root node.
   - Recursively, it constructs the left subtree (`start` to `mid-1`) and the right subtree (`mid+1` to `end`) by calling itself.
   - It returns the root of the balanced BST.

### Usage

To use this package:

1. Define a BST using `TreeNode` structs.
2. Pass the root of the BST to the `balanceBST` function.
3. The function will return the root of a balanced BST where the nodes are arranged such that the tree is as balanced as possible.

