func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	cache := make(map[[2]int]int)

	var bt func(int, int) int
	bt = func(acc, idx int) int {
		if idx >= n {
			if acc == target {
				return 1
			}
			return 0
		}

		key := [2]int{acc, idx}

		if _, ok := cache[key]; !ok {
			result := bt(acc + nums[idx], idx + 1) + bt(acc - nums[idx], idx + 1)
			cache[key] = result
		}

		return cache[key]
	}

	
	return bt(0, 0)
}
