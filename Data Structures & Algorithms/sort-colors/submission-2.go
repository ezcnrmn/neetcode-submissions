func sortColors(nums []int) {
	left, i, right := 0, 0, len(nums)-1
	for nums[left] == 0 {
		left++
		i = left
	}
	for nums[right] == 2 {
		right--
	}

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