func sortColors(nums []int) {
	left, i, right := 0, 0, len(nums)-1

	for i <= right {
		if nums[i] == 2 {
			nums[right], nums[i] = nums[i], nums[right]
			right--
			continue
		} else if nums[i] == 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
		i++
	}
}