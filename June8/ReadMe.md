#  Continuous SubArray Sum: June 8

## Intuition
The code aims to check whether there exists a contiguous subarray in the given array `nums` whose sum is divisible by `k`.

## Approach
The approach utilizes a map `sumMap` to store the cumulative sum modulo `k` along with their indices. This helps in identifying if a subarray with a sum divisible by `k` exists. If the difference between the current index and the index stored in `sumMap` for a particular cumulative sum is greater than 1, it indicates the presence of a subarray with a sum divisible by `k`.

## Complexity
- Time complexity: O(n), where n is the length of the input array `nums`. The algorithm iterates through the array once.
- Space complexity: O(n), where n is the length of the input array `nums`. The space used by the `sumMap` map is proportional to the number of elements in `nums`.
