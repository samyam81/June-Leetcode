# Sort Color : June 12

## Intuition
The Dutch National Flag problem involves sorting an array containing three different types of elements (in this case, 0s, 1s, and 2s) in linear time without using any extra space.

## Approach
We'll use a three-way partitioning algorithm to solve this problem. We'll maintain three pointers:
- `left` points to the next position where a 0 should be placed.
- `mid` scans the array from left to right.
- `right` points to the next position where a 2 should be placed.

We iterate through the array using the `mid` pointer. When we encounter a 0 at `nums[mid]`, we swap it with the element at `nums[left]` and increment both `left` and `mid`. When we encounter a 1, we simply increment `mid`. When we encounter a 2 at `nums[mid]`, we swap it with the element at `nums[right]` and decrement `right`. We repeat this process until `mid` crosses `right`.

## Complexity
- Time complexity: 
  - Since we only iterate through the array once, the time complexity is O(n), where n is the number of elements in the array.
- Space complexity:
  - We are using only a constant amount of extra space for storing variables and temporary values. Therefore, the space complexity is O(1).
