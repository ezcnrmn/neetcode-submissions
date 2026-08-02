func search(nums []int, target int) int {
	var left int
	right := len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > nums[left] {
			if target > nums[mid] || target <= nums[right] {
				left = mid+1
			} else {
				right = mid-1
			}
		} else {
			if target < nums[mid] || target >= nums[left] {
				right = mid-1
			} else {
				left = mid+1
			}
		}
	}
	return -1
}
