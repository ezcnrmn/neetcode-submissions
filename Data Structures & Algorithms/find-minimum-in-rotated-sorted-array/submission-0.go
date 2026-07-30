func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	result := 1001
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			result = min(result, nums[mid])
			right = mid-1
		} else {
			result = min(result, nums[right])
			left = mid+1
		}
	}
	return result
}
