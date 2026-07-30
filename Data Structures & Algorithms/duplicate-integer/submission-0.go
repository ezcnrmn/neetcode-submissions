func hasDuplicate(nums []int) bool {
	mem := make(map[int]struct{}) 
	for _, n := range nums {
		if _, ok := mem[n]; ok {
			return true
		}
		mem[n] = struct{}{}
	}
	return false
}
