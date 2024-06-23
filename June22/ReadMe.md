# Nice Sub-Array : June 22

### Intuition
The goal of this code is to find the number of subarrays in an integer array `nums` where the number of odd integers (elements with value `1` when performing a bitwise AND operation) is exactly `k`.

### Approach
The approach uses a prefix sum technique combined with an array `cnt` to keep track of the count of occurrences of different counts of odd integers encountered so far (`t`). Here's a step-by-step breakdown of the code:

1. **Initialization**:
   - `n` is the length of the `nums` array.
   - `cnt` is an array of size `n + 1` initialized to zero. This array will store counts of different values of `t` (the cumulative count of odd integers encountered so far).
   - `cnt[0]` is initialized to `1` because initially, there are no odd integers, and this represents the base case where `t` is zero.

2. **Iterating through `nums`**:
   - For each integer `v` in `nums`:
     - Update `t` by adding `v & 1`. This operation checks if `v` is odd (`v & 1` is `1`) or even (`v & 1` is `0`).
     - Check if the current `t - k` (where `k` is the desired number of odd integers in the subarray) is non-negative.
     - Add `cnt[t - k]` to `ans`. This value represents the number of subarrays ending at the current index of `nums` that satisfy the condition of having exactly `k` odd integers.

3. **Update `cnt`**:
   - Increment `cnt[t]` to record the occurrence of the current count `t` of odd integers encountered.

4. **Return Result**:
   - After processing all elements of `nums`, `ans` contains the total number of subarrays where the number of odd integers equals `k`.

### Complexity
- **Time complexity**: The algorithm runs in `O(n)` time, where `(n)` is the length of the `nums` array. This is because each element is processed exactly once.
- **Space complexity**: The space complexity is `O(n)` due to the `cnt` array used to store counts of different values of `t`.
