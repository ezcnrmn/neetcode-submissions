func combinationSum(nums []int, target int) [][]int {
	var result [][]int
	var cur []int

	var bt func(int, int)
	bt = func(sum, idx int) {
		if sum > target {
			return
		}
		if sum == target {
			cp := make([]int, len(cur))
			copy(cp, cur)
			result = append(result, cp)
			return
		}

		for i := idx; i < len(nums); i++ {
			cur = append(cur, nums[i])
			bt(sum+nums[i], i)
			cur = cur[:len(cur)-1]
		}
	}

	bt(0, 0)
	return result
}
