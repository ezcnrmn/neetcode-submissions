func sortColors(nums []int) {
	var counter [3]int
	for _, number := range nums {
		counter[number]++
	}
	var cur int
	for i := range nums {
		if counter[cur] == 0 {
			for counter[cur] == 0 {
				cur++
			}
		}
		nums[i] = cur
		counter[cur]--
	}
}
