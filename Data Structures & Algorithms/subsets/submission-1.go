func subsets(nums []int) [][]int {
	queue := [][]int{{}}
	for _, num := range nums {
		length := len(queue)
		for range length{
			cur := queue[0]
			queue = queue[1:]
			cur = append(cur, num)
			queue = append(queue, cur, cur[:len(cur)-1])
		}
	}
	return queue
}
