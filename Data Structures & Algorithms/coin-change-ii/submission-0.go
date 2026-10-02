func change(amount int, coins []int) int {
	if amount == 0 {
		return 0
	}
	
	n := len(coins)
	var result int

	var dfs func(int, int)
	dfs = func(index, sum int) {
		if sum > amount {
			return
		}
		if sum == amount {
			result++
			return
		}

		for i := index; i < n; i++ {
			dfs(i, sum+coins[i])
		}
	}

	dfs(0, 0)
	return result
}
