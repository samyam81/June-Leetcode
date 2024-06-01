# Score of String : June 1

## Intuition
At first glance, this solution seems to calculate the score of a given string based on the absolute difference between adjacent characters. 

## Approach
The `scoreOfString` method takes a string `s` as input and initializes a variable `result` to keep track of the score. Then, it iterates through the characters of the string using a loop, starting from the first character (index 0) up to the second to last character (`s.length() - 1`). Within each iteration, it calculates the absolute difference between the ASCII values of the current character (`s.charAt(i)`) and the next character (`s.charAt(i + 1)`), and adds this difference to the `result` variable. Finally, it returns the `result` as the score of the string.

## Complexity
- Time complexity: The time complexity of this solution is O(n), where n is the length of the input string `s`. This is because the solution iterates through each character of the string once.
- Space complexity: The space complexity of this solution is O(1) because it uses a constant amount of extra space, regardless of the size of the input string. The only extra space used is for the `result` variable, which stores the final score.
