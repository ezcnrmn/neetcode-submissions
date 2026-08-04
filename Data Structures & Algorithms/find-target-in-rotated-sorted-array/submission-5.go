func search(nums []int, target int) int {
	// 2 binary search solution
	left, right := 0, len(nums)-1
	var pivot int
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= nums[(mid+1)%len(nums)] {
			pivot = mid
			break
		}
		if nums[mid] > nums[left] {
			left = mid+1
		} else {
			right = mid-1
		}
	}

	if target <= nums[pivot] && target >= nums[0] {
		left = 0
		right = pivot
	} else if target >= nums[(pivot+1)%len(nums)] && target <= nums[len(nums)-1] {
		left = pivot+1
		right = len(nums)-1
	} else {
		return -1
	}

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid-1
		} else {
			left = mid+1
		}
	}
	return -1
}
