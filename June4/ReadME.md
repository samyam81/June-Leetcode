# Longest Palindrome : June 4

## Intuition
The code seems to be aimed at finding the length of the longest palindrome that can be formed using characters from a given string. It appears to utilize the frequency of characters to determine the palindrome's length.

## Approach
The code begins by creating a HashMap to store the frequency of each character in the input string. Then, it iterates through the HashMap to calculate the length of the longest palindrome that can be formed. It considers both even and odd frequencies of characters to ensure the formation of a valid palindrome.

## Complexity
- Time complexity:
  - The code iterates through the input string once to populate the HashMap, which takes O(n) time, where n is the length of the input string.
  - It then iterates through the keys of the HashMap, which takes O(m) time, where m is the number of unique characters in the input string.
  - Thus, the overall time complexity is O(n + m).

- Space complexity:
  - The code utilizes a HashMap to store the frequency of characters, which requires space proportional to the number of unique characters in the input string, i.e., O(m).
  - Additionally, it uses a few variables for bookkeeping, which require constant space.
  - Therefore, the overall space complexity is O(m), where m is the number of unique characters in the input string.
