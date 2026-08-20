func combinationSum(nums []int, target int) [][]int {
	queue := make([][]int, 1)
	queue[0] = []int{}
	for _, num := range nums {
		length := len(queue)
		for range length {
			cur := queue[0]
			queue = queue[1:]
			queue = append(queue, cur)
			cur2 := make([]int, len(cur))
			copy(cur2, cur)
			for {
				if sum(cur2) + num <= target {
					cur3 := make([]int, len(cur2)+1)
					copy(cur3, cur2)
					cur3[len(cur3)-1] = num
					queue =append(queue, cur3)
					cur2 = cur3
				} else {
					break
				}
			}
		}
	}

	var result [][]int
	for _, q := range queue {
		if sum(q) == target {
			result = append(result, q)
		}
	}
	return result
}

func sum(nums []int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}