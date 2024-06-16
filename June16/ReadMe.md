# Patching Array : June 16

## Intuition
The problem requires us to determine the minimum number of patches needed to cover all numbers from 1 to n using elements from a sorted array nums. If nums cannot provide enough numbers to cover all gaps up to n, we must strategically add patches to fill these gaps efficiently.

## Approach
We maintain a variable `miss` to track the smallest number that cannot be formed using the current set of patches. We iterate through `nums` and use each number to extend the range that can be formed up to `miss + nums[i]`. If `nums[i]` is larger than `miss` or `nums` is exhausted (i exceeds bounds), we add `miss` itself as a patch and double `miss` to extend the range. We repeat this until `miss` exceeds `n`, counting each added patch.

## Complexity

** Time complexity: **
The main part of the algorithm runs in O(m + log n)time, where (m) is the number of elements in `nums` and (n) is the given integer up to which we need to cover the range. This complexity arises from iterating through `nums` and potentially doubling `miss` logarithmically relative to `n`.

** Space complexity: **
The space complexity is O(1) because the algorithm uses only a constant amount of extra space regardless of the size of `nums` or `n`. It primarily operates with a few variables (`miss`, `result`, and `i`) to keep track of the current state and count patches. Thus, the space used is constant and does not depend on the input size.
