func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	var result int

	var bt func(int, int)
	bt = func(acc, idx int) {
		if idx >= n {
			if acc == target {
				result++
			}
			return
		}

		bt(acc + nums[idx], idx + 1)
		bt(acc - nums[idx], idx + 1)
	}

	bt(0, 0)
	return result
}
