func maxSubArray(nums []int) int {
	var cur int
	result := nums[0]
	for _, num := range nums {
		cur += num
		result = max(result, cur)
		if cur < 0 {
			cur = 0
		}
	}
	return result
}
