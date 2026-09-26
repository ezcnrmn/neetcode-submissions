func permute(nums []int) [][]int {
	var result [][]int
	
	var bt func([]int, map[int]struct{})
	bt = func(acc []int, used map[int]struct{}) {
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

				bt(acc, used)

				delete(used, n)
				acc = acc[:len(acc)-1]
			}
		}
	}

	initial := make([]int, 0, len(nums))
	used := make(map[int]struct{}, len(nums))
	bt(initial, used)

	return result
}
