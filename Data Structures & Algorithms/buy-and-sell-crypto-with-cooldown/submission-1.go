func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	n := len(prices)
	dpCool := 0
	dpWait := 0
	dpHold := -prices[0]

	for i := 2; i <= n; i++ {
		dpCool, dpWait, dpHold = dpHold + prices[i-1], max(dpWait, dpCool), max(dpHold, dpWait-prices[i-1])
	}

	return max(dpCool, dpWait, dpHold)
}
