# Minimum Increment to Make Array Unique : June 14

## Intuition
The primary thought was to sort the array and then iterate through it. If an element is found to be equal to or less than its previous element, increment it to make it unique while keeping track of the total number of increments needed.

## Approach
1. Sort the array `nums`.
2. Initialize a variable `tracker` to keep track of the total increments needed.
3. Iterate through the sorted array starting from the second element.
4. If the current element is less than or equal to the previous element, calculate the difference between them and increment the current element to make it unique. Add the required number of increments to the `tracker`.
5. Continue this process until the end of the array is reached.
6. Return the value of `tracker` as the result.

## Complexity
- Time complexity: O(n log n), where n is the number of elements in the array. This is due to the sorting operation.
- Space complexity: O(1), as the space used is independent of the input size, except for the input itself.
