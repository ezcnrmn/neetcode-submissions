func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	reach := make([]int, n+1)
	for i := 2; i <= n; i++ {
		step1 := cost[i-1] + reach[i-1]
		step2 := cost[i-2] + reach[i-2]
		reach[i] = min(step1, step2)
	}
	return reach[n]
}
