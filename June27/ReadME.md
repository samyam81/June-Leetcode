# Find Center of Star Graph: June 27

## Intuition
To find the center of a tree represented by edges, we need to identify the node that has a degree (number of connected edges) equal to `n-1`, where `n` is the number of nodes in the tree. In a tree with `n` nodes, all nodes except the center will have a degree of exactly 1, and the center will have a degree of `n-1`.

## Approach
1. **Initialize**: Determine the number of nodes `n` in the tree, which is `len(edges) + 1` because there are `len(edges)` edges connecting `len(edges) + 1` nodes.
2. **Degree Counting**: Create an array `degreeCount` of size `n+1` to store the degree of each node. Traverse through each edge and increment the degree count for both nodes involved in the edge.
3. **Find the Center**: Traverse through the `degreeCount` array from 1 to `n`. The first node that has a degree equal to `n-1` is the center of the tree.

## Complexity
- Time complexity: **O(n)**, where `n` is the number of nodes in the tree. We iterate through the edges to compute degrees (`O(e)`, where `e` is the number of edges) and then iterate through the nodes (`O(n)`) to find the center.
- Space complexity: **O(n)**. We use additional space for the `degreeCount` array, which has `n+1` elements. Thus, the space complexity is linear in terms of the number of nodes.