func canJump(nums []int) bool {
	var i int
	for i < len(nums) {
		if i == len(nums)-1 {
			break
		}
		if nums[i] == 0 {
			return false
		}
		i += nums[i]
	}
	return true
}
