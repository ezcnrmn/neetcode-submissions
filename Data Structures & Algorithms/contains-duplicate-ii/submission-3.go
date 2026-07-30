func containsNearbyDuplicate(nums []int, k int) bool {
	window := make(map[int]struct{})
	for i := 1; i <= k; i++ {
		if _, ok := window[nums[i]]; ok {
			return true
		}
		window[nums[i]] = struct{}{}
	}
	for i := range len(nums)-k {
		if _, ok := window[nums[i]]; ok {
			return true
		}
		if i + k + 1 < len(nums) {
			delete(window, nums[i+1])
			window[nums[i+k+1]] = struct{}{}
		}
	}
	return false
}
