# Relative Sort Array: June 11

## Intuition
The code appears to be implementing a function to sort an array (`arr1`) based on the relative order of elements in another array (`arr2`). It counts the frequency of each element in `arr1`, then iterates over `arr2` to place elements in the result array according to their frequency. Finally, it fills the remaining elements in the result array with the elements from `arr1` that are not present in `arr2`, in ascending order.

## Approach
1. Initialize a frequency array to count the occurrences of each element in `arr1`.
2. Iterate over `arr1` and update the frequency array accordingly.
3. Initialize a result array to store the sorted elements.
4. Iterate over `arr2`. For each element in `arr2`, append it to the result array as many times as its frequency indicates.
5. Iterate over the frequency array. For each non-zero frequency, append the corresponding element to the result array as many times as its frequency indicates.
6. Return the result array.

## Complexity
- Time complexity:
  - Counting the frequency of elements in `arr1` takes linear time, O(n), where n is the length of `arr1`.
  - Iterating over `arr2` and `frequency` takes linear time as well, O(m), where m is the length of `arr2`.
  - Overall, the time complexity is O(n + m), where n is the length of `arr1` and m is the length of `arr2`.

- Space complexity:
  - The space complexity is O(1) for the frequency array since it has a fixed size of 1001.
  - The space complexity of the result array is O(n), where n is the length of `arr1`.
  - Therefore, the overall space complexity is O(n) due to the result array.
