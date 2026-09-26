func permute(nums []int) [][]int {
	var result [][]int

	acc := make([]int, 0, len(nums))
	used := make(map[int]struct{}, len(nums))
	
	var bt func()
	bt = func() {
		if len(acc) == len(nums) {
			cp := make([]int, len(nums))
			copy(cp, acc)
			result = append(result, cp)
			return
		}

		for _, n := range nums {
			if _, wasUsed := used[n]; !wasUsed {
				used[n] = struct{}{}
				acc = append(acc, n)

				bt()

				delete(used, n)
				acc = acc[:len(acc)-1]
			}
		}
	}

	bt()

	return result
}
