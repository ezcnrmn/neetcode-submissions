func majorityElement(nums []int) int {
	res, amount := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == res {
			amount++
		} else if amount == 0 {
			res = nums[i]
			amount = 1
		} else {
			amount--
		}
	}
	return res
}
