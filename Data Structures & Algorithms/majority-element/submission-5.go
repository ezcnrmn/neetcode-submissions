func majorityElement(nums []int) int {
	var cur, counter int
	for _, num := range nums {
		if counter == 0 {
			cur = num
			counter = 1
			continue
		}
		if cur == num {
			counter++
		} else {
			counter--
		}
	}
	return cur
}
