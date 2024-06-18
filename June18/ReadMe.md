# Most Profit Assigning Work : June 18


## Intuition
The problem requires assigning maximum profits to workers based on their skill levels and job difficulties. Each worker can handle jobs up to a certain difficulty level, and the objective is to maximize the total profit across all workers.

## Approach
1. **Find Maximum Difficulty Level**: Determine the highest difficulty level among all jobs to establish the range for storing maximum profits.

2. **Initialize Data Structures**: Use an array `maxProfitUpToDifficulty` where each index represents a difficulty level, and the value at each index holds the maximum profit achievable up to that difficulty level.

3. **Compute Maximum Profits**: Populate `maxProfitUpToDifficulty` by iterating through the list of jobs. For each job, update the maximum profit at its difficulty level if a higher profit is found.

4. **Fill in Missing Values**: Ensure that `maxProfitUpToDifficulty` is filled correctly by iterating through the array and setting each value to the maximum of itself and the previous difficulty level's maximum profit.

5. **Calculate Total Profit**: Iterate through each worker's skill level. If a worker's skill level exceeds the maximum difficulty level, assign them the maximum profit achievable at the highest difficulty level. Otherwise, assign them the maximum profit corresponding to their skill level.

6. **Return Total Profit**: Accumulate and return the total profit across all workers.

## Complexity
- Time complexity:
  - The primary operations involve iterating through the lists of jobs and workers, making the time complexity primarily dependent on these inputs. Hence, the time complexity is **O(D + N + W)**, where D is the number of jobs (and thus the maximum difficulty level), N is the total number of jobs, and W is the number of workers.
  
- Space complexity:
  - The space complexity is **O(D)**, where D is the maximum difficulty level. This is because we use an array `maxProfitUpToDifficulty` of size `D+1` to store maximum profits up to each difficulty level.
