func longestConsecutive(nums []int) int {
	mem := make(map[int]struct{}, len(nums))
	for _, number := range nums {
		mem[number] = struct{}{}
	}

	var result int
	for number := range mem {
		segment := [2]int{number, number}
		delete(mem, number)
		for {
			if _, ok := mem[segment[0]-1]; !ok {
				break
			}
			segment[0]--
			delete(mem, segment[0])
		}
		for {
			if _, ok := mem[segment[1]+1]; !ok {
				break
			}
			segment[1]++
			delete(mem, segment[1])
		}
		result = max(result, segment[1]-segment[0]+1)
	}

	return result
}