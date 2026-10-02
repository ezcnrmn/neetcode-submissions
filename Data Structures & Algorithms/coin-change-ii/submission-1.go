func change(amount int, coins []int) int {
	if amount == 0 {
		return 0
	}

	n := len(coins)
	cache := make(map[[2]int]int)

	var dfs func(int, int) int
	dfs = func(index, sum int) int {
		if sum > amount {
			return 0
		}
		if sum == amount {
			return 1
		}

		if val, ok := cache[[2]int{index, sum}]; ok {
			return val
		}

		var cur int
		for i := index; i < n; i++ {
			cur += dfs(i, sum+coins[i])
		}

		cache[[2]int{index, sum}] = cur
		return cur
	}

	return dfs(0, 0)
}
