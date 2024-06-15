# IPO : June 15

## Intuition
The problem requires selecting projects to maximize capital while adhering to capital constraints. We need an efficient strategy to ensure that we maximize profit while keeping track of available capital.

## Approach
1. **Data Representation**: We represent each project with its capital requirement and profit in a struct.
2. **Sorting**: First, we sort all projects based on their capital requirements. This allows us to consider projects in increasing order of their capital.
3. **Max-Heap for Profit**: We use a max-heap (implemented with a min-heap and inverting profits) to efficiently select the project with the maximum profit that we can afford at any point.
   - We iterate through sorted projects and push their profits into the heap if they are affordable with current capital.
   - Each time we need to select a project to invest in, we pop the project with the highest profit from the heap.
4. **Iteration**: We iterate through this process until we've selected `k` projects or cannot afford any more projects.

This approach ensures that we always have the option to pick the project with the maximum profit that fits within our current available capital.

## Complexity
- Time complexity: Sorting the projects takes \( O(n \log n) \). Each project is pushed and popped from the heap at most once, resulting in \( O(n \log k) \) operations due to the heap operations. Thus, the overall time complexity is dominated by the sorting step, \( O(n \log n) \).
- Space complexity: We use additional space for storing the projects and the heap. The space complexity is \( O(n) \) for storing projects and \( O(k) \) for the heap, resulting in \( O(n) \) space complexity overall.
