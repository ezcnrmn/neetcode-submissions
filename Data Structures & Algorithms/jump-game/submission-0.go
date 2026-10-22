func canJump(nums []int) bool {
	var i int
	for i < len(nums) {
		if i == len(nums)-1 {
			return true
		}
		if nums[i] == 0 {
			break
		}
		i += nums[i]
	}
	return false
}
