# Grumpy Bookstore Owner : June 21

## Intuition

We want to maximize the number of satisfied customers by utilizing the "grumpy" technique only when it benefits us the most.

## Approach

1. Calculate the initial total satisfied customers without applying the grumpy technique.
2. Use a sliding window approach to find the maximum additional satisfied customers we can get by applying the grumpy technique in a contiguous window of 'minutes' length.
3. Sum up the initial satisfied customers and the additional satisfied customers obtained from the window to get the result.

## Complexity
- Time complexity:

   The time complexity of the algorithm is O(n), where n is the length of the customers array. This is because we make a single pass through the customers and grumpy arrays, and the sliding window operates in linear time.
   
- Space complexity:

   The space complexity is O(1) since we are using only a constant amount of extra space regardless of the input size.


