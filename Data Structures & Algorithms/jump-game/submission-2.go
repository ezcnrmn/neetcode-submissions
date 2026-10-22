func canJump(nums []int) bool {
	var balance int
	for i := 0; i < len(nums)-1; i++ {
		balance = max(balance-1, nums[i])
		if balance == 0 {
			return false
		}
	}
	return true
}
