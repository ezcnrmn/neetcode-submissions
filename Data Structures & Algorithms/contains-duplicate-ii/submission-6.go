func containsNearbyDuplicate(nums []int, k int) bool {
	left, right := 0, 0
	window := make(map[int]struct{})
	window[nums[0]] = struct{}{}
	for left <= len(nums)-k {
		if right - left == k {
			delete(window, nums[left])
			left++
		}

		if right < len(nums)-1 {
			right++

			if _, ok := window[nums[right]]; ok { return true }
			window[nums[right]] = struct{}{}
		} 
	}
	return false
}
