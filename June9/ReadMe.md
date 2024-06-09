# Subarray Sums Divisible by K: June 9

## Intuition
This code aims to count the number of subarrays whose sum is divisible by a given integer `k`.

## Approach
The code utilizes the concept of prefix sum and remainder to efficiently count the subarrays. It maintains a `prefixsum` variable to keep track of the cumulative sum of the elements encountered so far. It also uses a `remainder` map to store the frequency of remainders when dividing the prefix sum by `k`. 

## Complexity
- Time complexity: O(n), where n is the length of the input array `nums`. The code iterates through each element of the array once.
- Space complexity: O(k), where k is the value of the input integer `k`. The space complexity arises from the `remainder` map, which can store at most k unique remainders.
