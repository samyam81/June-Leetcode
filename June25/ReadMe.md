# BST to Greater Sum Tree (Go Solution)

This Go package (`go-solution`) provides a solution for transforming a Binary Search Tree (BST) into a Greater Sum Tree (GST). The transformation is achieved by modifying each node's value to be the sum of all nodes greater than itself in the original BST.

## TreeNode Struct

The package defines a `TreeNode` struct which represents a node in the BST:

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
```

## Function: bstToGst

The main function provided by this package is `bstToGst`, which takes a root node of a BST as input and returns the root node of the corresponding GST.

### Explanation

The `bstToGst` function uses a depth-first search (DFS) approach to convert the BST into a GST:

- **DFS Function (`dfs`)**: This function recursively traverses the BST in a specific order (`Right`, `Root`, `Left`). It accumulates the sum of all greater nodes in the `sum` variable and updates each node's value accordingly.

- **Update Process**:
  - Traverse the right subtree first to accumulate the sum of all greater nodes.
  - Update the current node's value (`node.Val`) to be the accumulated sum.
  - Traverse the left subtree to continue the process recursively.

### Usage

To use this package:

1. Define a BST using `TreeNode` structs.
2. Pass the root of the BST to the `bstToGst` function.
3. The function will modify the BST in-place to convert it into a GST.
4. Use the modified root returned by `bstToGst` for further processing or evaluation.

For more details, refer to the function's implementation in the code.
