func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	n := len(prices)
	dp_cool := make([]int, n+1)
	dp_wait := make([]int, n+1)
	dp_hold := make([]int, n+1)

	dp_hold[1] = -prices[0]

	for i := 2; i <= n; i++ {
		dp_cool[i] = dp_hold[i-1] + prices[i-1]
		dp_wait[i] = max(dp_wait[i-1], dp_cool[i-1])
		dp_hold[i] = max(dp_hold[i-1], dp_wait[i-1]-prices[i-1])
	}

	fmt.Println(dp_cool)
	fmt.Println(dp_wait)
	fmt.Println(dp_hold)

	return max(dp_cool[n], dp_wait[n], dp_hold[n])
}
