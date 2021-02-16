func coinChange(coins []int, amount int) int {
	var dp = make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		res := math.MaxInt32
		for _, coin := range coins {
			if i >= coin {
				res = min(res, dp[i-coin])
			}
		}
		dp[i] = min(dp[i], res+1)
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	} else {
		return dp[amount]
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
