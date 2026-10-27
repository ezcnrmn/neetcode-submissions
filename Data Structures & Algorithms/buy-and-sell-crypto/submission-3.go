func maxProfit(prices []int) int {
	var result, maxPrice int
	for i := len(prices)-1; i >= 0; i-- {
		maxPrice = max(maxPrice, prices[i])
		diff := maxPrice-prices[i]
		result = max(result, diff)
	}
	return result
}
