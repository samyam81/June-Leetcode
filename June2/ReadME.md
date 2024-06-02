# Reverse The String : June 2

## Intuition
Reversing a string can be efficiently done by swapping characters from both ends until they meet in the middle.

## Approach
 We initialize two pointers, one pointing to the start of the string and the other pointing to the end. We swap the characters at these pointers iteratively until they converge.

## Complexity
- Time complexity:
The time complexity is O(n), where n is the length of the input string. We only need to iterate through half of the string to perform the swapping operation.

- Space complexity:
The space complexity is O(1) since we are not using any extra space that grows with the input size. 