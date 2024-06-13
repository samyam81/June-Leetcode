# Minimum Number of Moves to Seat Everyone : June 13

## Intuition
To minimize the number of moves required to seat students, we aim to match each student to the closest available seat. 

## Approach
1. Sort both the `seats` and `students` slices to ensure they are in ascending order.
2. Iterate through both slices simultaneously.
3. For each pair of corresponding elements (seat and student), calculate the absolute difference to determine the distance between the seat and the student.
4. Accumulate the total moves by adding up these absolute differences.
5. Return the total moves.

## Complexity
- Time complexity: O(n log n) due to sorting the `seats` and `students` slices.
- Space complexity: O(1) since we only use a constant amount of extra space for variables regardless of the input size.
