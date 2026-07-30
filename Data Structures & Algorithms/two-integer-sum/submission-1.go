func twoSum(nums []int, target int) []int {
	mem := make(map[int]int, len(nums))
	for i, v := range nums {
		mem[v] = i
	}
	for i, v := range nums {
		if j, ok := mem[target-v]; ok && j != i {
			return []int{i, j}
		}
	}
	return []int{0, 0}
}