func findDuplicate(nums []int) int {
	slow, fast, prev := 0, 0, 0
	for {
		prev = nums[fast]
		fast = nums[fast]
		if slow == fast {
			return prev
		}
		prev = nums[fast]
		fast = nums[fast]
		if slow == fast {
			return prev
		}
		slow = nums[slow]
	}
	return -1
}
