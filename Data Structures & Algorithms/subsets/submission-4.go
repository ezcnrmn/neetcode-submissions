func subsets(nums []int) [][]int {
	n := len(nums)
	var result [][]int

	acc := make([]int, 0, n)

	var bt func(int)
	bt = func(idx int) {
		cp := make([]int, len(acc))
		copy(cp, acc)
		result = append(result, cp)

		for i := idx; i < n; i++ {
			acc = append(acc, nums[i])
			bt(i+1)
			acc = acc[:len(acc)-1]
		}
	}

	bt(0)
	return result
}
