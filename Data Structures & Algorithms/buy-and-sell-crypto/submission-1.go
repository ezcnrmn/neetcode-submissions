func maxProfit(prices []int) int {
	var left, result int
	for right := 0; right < len(prices); right++ {
		profit := prices[right]-prices[left]
		if profit < 0 {
			left = right
		} else {
			result = max(result, profit)
		}
	}
	return result
}
