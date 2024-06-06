# Hand of Straights: June 6

## Intuition
We need to check if it's possible to rearrange the given hand of cards into groups such that each group consists of `groupSize` consecutive cards. To solve this, we can iterate through the cards and try to form groups starting from each card.

## Approach
1. First, we check if the length of the hand is divisible by `groupSize`. If it's not, it's impossible to form groups of the desired size, so we return false.
2. We create a map to store the count of each card value in the hand.
3. We sort the hand to make it easier to iterate through consecutive cards.
4. We iterate through each card in the sorted hand.
   - For each card, we check if there are enough remaining cards to form a group starting from that card.
   - If there are enough cards, we decrement their counts in the map.
   - If at any point we encounter a card with a count of 0 in the map, it means there aren't enough cards to form a group, so we return false.
5. If we successfully form groups for all cards, we return true.

## Complexity
- Time complexity:  
  - Sorting the hand takes O(n log n) time.
  - The iteration through the hand takes O(n) time.
  - Overall, the time complexity is O(n log n).
  
- Space complexity:
  - We use a map to store the count of each card value, which takes O(n) space.
  - Additionally, sorting the hand takes O(n) space.
  - Therefore, the space complexity is O(n).
