import "slices"

func threeSum(nums []int) [][]int {
	set := make(map[[3]int]struct{})
	cur := make([]int, 0, 3)
	
	var dfs func(int)
	dfs = func(idx int) {
		if len(cur) == 3 {
			if cur[0] + cur[1] + cur[2] == 0 {
				var cp [3]int
				copy(cp[:], cur)
				slices.Sort(cp[:])
				set[cp] = struct{}{}
			}
			return
		}

		for i := idx; i < len(nums); i++ {
			cur = append(cur, nums[i])
			dfs(i+1)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(0)

	result := make([][]int, 0, len(set))
	for s := range set {
		result = append(result, s[:])
	}
	return result
}
