# Sum of Square Number : June 17

### Intuition
The problem requires determining whether there exist two non-negative integers \( a \) and \( b \) such that \( a^2 + b^2 = c \). Given the nature of the problem, a direct iterative approach might be inefficient for large values of \( c \). Instead, leveraging mathematical properties of squares can lead to a more efficient solution.

### Approach
The chosen approach employs a two-pointer technique combined with mathematical operations:
- **Initialization**: Start with `left` initialized to 0 and `right` initialized to \( \sqrt{c} \) (converted to an integer). This choice of `right` is based on the maximum possible value for \( b \).
- **Two-pointer Technique**: Use a loop where `left` and `right` pointers move inward towards each other:
  - Calculate `sum` as \( \text{left}^2 + \text{right}^2 \).
  - If `sum` equals `c`, return `true` as we found \( a \) and \( b \) such that \( a^2 + b^2 = c \).
  - If `sum` is less than `c`, increment `left` (`left++`).
  - If `sum` is greater than `c`, decrement `right` (`right--`).
- **Termination**: If `left` exceeds `right`, exit the loop and return `false`, indicating no valid pair \( (a, b) \) exists such that \( a^2 + b^2 = c \).

### Complexity
- **Time Complexity**: \( O(\sqrt{c}) \)
  - The main loop runs at most \( \sqrt{c} \) times because `left` starts from 0 and `right` starts from \( \sqrt{c} \). Each iteration adjusts either `left` or `right`, efficiently reducing the search space.
- **Space Complexity**: \( O(1) \)
  - The algorithm uses only a constant amount of extra space regardless of the input size \( c \). This includes variables like `left`, `right`, and `sum`.

This approach ensures an efficient solution in \( O(\sqrt{c}) \) time, suitable for large values of \( c \), with \( O(1) \) space complexity.

