# Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit : June 23

## Intuition
The problem requires finding the longest subarray such that the absolute difference between the maximum and minimum element in that subarray is less than or equal to a given limit.

## Approach
To solve this problem efficiently, we can use a sliding window approach with two deques (double-ended queues):
- `maxq`: This deque keeps track of the maximum elements in the current window.
- `minq`: This deque keeps track of the minimum elements in the current window.

Here's a step-by-step explanation of the code:

1. **Initialization**:
   - `maxq` and `minq` are initialized as empty deques.
   - `j` is set to 0 to mark the start of the current window.
   - `ans` is initialized to 0 to store the maximum length of valid subarrays found.

2. **Sliding Window Iteration**:
   - Iterate through each element `nums[i]` in the array:
     - **Maintain `maxq`**: Remove elements from `maxq` that are less than the current element `nums[i]`, ensuring `maxq` remains sorted in descending order.
     - **Maintain `minq`**: Remove elements from `minq` that are greater than the current element `nums[i]`, ensuring `minq` remains sorted in ascending order.
     - Add `nums[i]` to both `maxq` and `minq`.

   - **Adjust Window**: While the difference between the maximum (`maxq[0]`) and minimum (`minq[0]`) elements in the current window exceeds `limit`, increment `j` (move the window start) and adjust `maxq` and `minq` accordingly.

   - **Update Answer**: For each `i`, calculate the length of the current valid subarray (`i - j + 1`) and update `ans` with the maximum length found so far.

3. **Return Result**: After iterating through the entire array, `ans` contains the length of the longest subarray satisfying the given condition.

## Complexity
- **Time complexity**: The algorithm runs in  **O(n)** time, where `n` is the length of the `nums` array. This is because each element is processed a constant number of times across all iterations.
- **Space complexity**: The space complexity is *O(n)*  due to the storage used by the deques `maxq` and `minq`.
