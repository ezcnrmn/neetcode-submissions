func subsets(nums []int) [][]int {
	queue := [][]int{{}}
	for _, num := range nums {
		length := len(queue)
		for range length{
			cur := queue[0]
			queue = queue[1:]
			next := make([]int, len(cur)+1)
			copy(next, cur)
			next[len(next)-1] = num
			queue = append(queue, cur, next)
		}
	}
	return queue
}
