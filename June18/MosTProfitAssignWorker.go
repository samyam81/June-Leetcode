package main

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
    var maxdifficulty int = 0
    for _, d := range difficulty {
        maxdifficulty = max(maxdifficulty, d)
    }

    maxProfitUpToDifficulty := make([]int, maxdifficulty+1)

    for i := 0; i < len(difficulty); i++ {
        maxProfitUpToDifficulty[difficulty[i]] = max(maxProfitUpToDifficulty[difficulty[i]], profit[i])
    }

    for i := 1; i <= maxdifficulty; i++ {
        maxProfitUpToDifficulty[i] = max(maxProfitUpToDifficulty[i], maxProfitUpToDifficulty[i-1])
    }

    var totalProfit int = 0
    for _, caliber := range worker {
        if caliber > maxdifficulty {
            totalProfit += maxProfitUpToDifficulty[maxdifficulty]
        } else {
            totalProfit += maxProfitUpToDifficulty[caliber]
        }
    }

    return totalProfit
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
