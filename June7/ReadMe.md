#  Replace Words : June 7

## Intuition
The code presents a function designed to replace words in a sentence with their shortest root word from a given dictionary. 

## Approach
1. **Split Sentence:** Split the input sentence into individual words.
2. **Create Root Map:** Generate a map of root words from the provided dictionary.
3. **Iterate Over Words:** Loop through each word in the sentence.
4. **Check Prefixes:** For each word, iterate over its prefixes.
5. **Replace Words:** If a prefix is found in the root map, replace the word with that prefix and break the inner loop.
6. **Join Modified Words:** Concatenate the modified words back into a sentence and return it.

## Complexity
- **Time Complexity:**
  - Let `n` be the number of words in the sentence and `m` be the maximum length of a word.
  - Building the root map takes `O(d)`, where `d` is the number of words in the dictionary.
  - Iterating over each word and its prefixes takes `O(n * m)`, where `n` is the number of words and `m` is the maximum length of a word.
  - Overall, the time complexity is `O(d + n * m)`.

- **Space Complexity:**
  - The space complexity is primarily determined by the root map, which takes `O(d)` space, where `d` is the number of words in the dictionary.
  - Additional space is used for storing intermediate variables like words, which takes `O(n)` space, where `n` is the number of words in the sentence.
  - Therefore, the space complexity is `O(d + n)`.
