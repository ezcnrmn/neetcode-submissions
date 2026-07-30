func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums)-k; i++ {
		for j := i+1; j <= i+k; j++ {
			if nums[i] == nums[k] {
				return true
			}
		}
	}
	return false
}
