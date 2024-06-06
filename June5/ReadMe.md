# Find Common Characters: June 5

## Intuition
To find the characters that appear in all strings within the input array, we can use a hashmap to count the characters in the first word, then iterate through the rest of the words to update the counts accordingly.

## Approach
1. Initialize a hashmap to store the count of characters in the first word.
2. Iterate through the characters in the first word and update the count in the hashmap.
3. Iterate through the rest of the words.
4. For each word, create a new hashmap to count the characters.
5. Update the counts in the main hashmap with the minimum counts from the current word.
6. Construct the result array based on the characters that appear in all words.

## Complexity
- Time complexity: O(n * m), where n is the number of words in the input array and m is the maximum length of a word. We iterate through each character in each word.
- Space complexity: O(k), where k is the number of unique characters that appear in all words. We store the character counts in the hashmap.
