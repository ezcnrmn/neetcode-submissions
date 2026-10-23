func jump(nums []int) int {
	result := len(nums)-1

	var bt func(int, int)
	bt = func(idx, jumps int) {
		if idx >= len(nums)-1 {
			result = min(result, jumps)
		}
		if jumps >= result {
			return
		}

		for i := idx+1; i < len(nums) && i <= idx+nums[idx]; i++ {
			bt(i, jumps+1)
		}
	}

	bt(0, 0)

	return result
}
