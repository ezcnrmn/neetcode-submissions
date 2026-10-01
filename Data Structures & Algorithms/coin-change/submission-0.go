func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		cur := math.MaxInt
		for _, coin := range coins {
			if coin <= i && dp[i-coin] != -1 {
				cur = min(cur, 1 + dp[i-coin])
			}
		}
		if cur != math.MaxInt {
			dp[i] = cur
		}
	}
	return dp[amount]
}
