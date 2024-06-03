# Append Characters to String to Make Subsequence: June 3

## Intuition
The most straightforward approach to solve this problem would be to iterate through both strings `s` and `t` and check if the characters match. If they match, we can move on to the next characters in both strings. If they don't match, then we need to keep track of how many characters in `t` we've "skipped" so far. These skipped characters would be the characters that need to be appended to `s` to make it a subsequence of `t`.

## Approach

The code implements the intuitive approach described above. It uses two pointers `i` and `j` to iterate through strings `s` and `t`, respectively. The key idea is to only increment `j` (pointer for `t`) when the characters at `i` and `j` match. This ensures that we are only considering characters in `t` that are part of the potential subsequence. By keeping track of the difference between the length of `t` and the value of `j` (number of characters matched), we can determine the number of characters that need to be appended to `s`.

## Complexity

- Time complexity:  The time complexity of the code is  **O(min(m, n))**, where `m` and `n` are the lengths of strings `s` and `t`, respectively. This is because the `while` loop iterates through the shorter of the two strings in the worst case. In each iteration, constant time operations are performed.

- Space complexity: The space complexity of the code is  **O(1)**. It only uses a constant amount of extra space for variables like `i`, `j`, `m`, and `n`.