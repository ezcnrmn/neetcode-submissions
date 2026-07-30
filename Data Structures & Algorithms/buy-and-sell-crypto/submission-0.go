func maxProfit(prices []int) int {
	var left, result int
	for right := 0; right < len(prices); right++ {
		window := prices[right]-prices[left]
		if window < 0 {
			left = right
			window = 0
		} else {
			result = max(result, window)
		}
	}
	return result
}
