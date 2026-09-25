type poolItem struct {
	num    int
	amount int
}

func combinationSum2(candidates []int, target int) [][]int {
	count := make(map[int]int, len(candidates))
	for _, c := range candidates {
		count[c]++
	}

	pool := make([]poolItem, 0, len(count))
	for num, amount := range count {
		pool = append(pool, poolItem{num: num, amount: amount})
	}

	var result [][]int

    var r func([]int, int) 
    r = func (cur []int, idx int) {
		s := sum(cur)
		if s == target {
			cp := make([]int, len(cur))
			copy(cp, cur)
			result = append(result, cp)
			return
		}
		if s > target {
			return
		}

        for i := idx; i < len(pool); i++ {
			for range pool[i].amount {
				cur = append(cur, pool[i].num)
				r(cur, i+1)
			}
			cur = cur[:len(cur)-pool[i].amount]
        }
    }

    r([]int{}, 0)

    return result
}

func sum(nums []int) (acc int) {
	for i := range nums {
		acc += nums[i]
	}
	return
}