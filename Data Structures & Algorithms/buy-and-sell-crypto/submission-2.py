class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        max_profit, min_price = 0, prices[0]
        for i in range(1, len(prices)):
            if prices[i] < min_price:
                min_price = prices[i]
            else:
                max_profit = max(max_profit, prices[i] - min_price)
        return max_profit
