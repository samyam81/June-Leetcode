# Height Checker : June 10

## Intuition
The code seems to be solving a problem related to checking the number of elements in an array that are not in their expected order.

## Approach
The code first creates a copy of the input array `heights` called `expected` and sorts it. Then, it iterates through both `heights` and `expected` arrays simultaneously, comparing each element. If an element in `heights` is not equal to its corresponding element in `expected`, it increments the `count`. Finally, it returns the count, which represents the number of elements that are not in their expected order.

## Complexity
- Time complexity:
  - Sorting the `expected` array takes O(n log n) time, where n is the length of the input array.
  - Iterating through both arrays takes O(n) time.
  - So, the overall time complexity is O(n log n).

- Space complexity:
  - Additional space is used for the `expected` array, which has the same length as the input array, so it takes O(n) space.
  - Apart from that, the space complexity is constant.
  - So, the overall space complexity is O(n).
