func mostPoints(questions [][]int) int64 {
    n := len(questions)
    dp := make([]int, n+1)

    for i := n - 1; i >= 0; i-- {
        solve := questions[i][0]
        next := min(n, i+questions[i][1]+1) 
        skip := dp[i+1]

        dp[i] = max(solve+dp[next], skip)
    }

    return int64(dp[0])
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}