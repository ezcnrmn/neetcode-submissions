func containsNearbyDuplicate(nums []int, k int) bool {
	left, right := 0, 0
	window := make(map[int]struct{})
	for right < len(nums) {
		if right - left > k {
			delete(window, nums[left])
			left++
		}

		if _, ok := window[nums[right]]; ok { return true }
		window[nums[right]] = struct{}{}
		right++
	}
	return false
}
