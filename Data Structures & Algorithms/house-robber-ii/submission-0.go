func rob(nums []int) int {
	cache := make(map[[2]int]int)
	n := len(nums)

	var bt func(int, int) int
	bt = func(from, to int) int {
		if from > to {
			return 0
		}
		if to == from {
			return nums[to]
		}

		if _, ok := cache[[2]int{from, to}]; !ok {
			cache[[2]int{from, to}] = max(nums[from] + bt(from+2, to), bt(from+1, to))
		}

		return cache[[2]int{from, to}]
	}

	result := max(nums[0] + bt(2, n-2), bt(1, n-1))
	return result
}
