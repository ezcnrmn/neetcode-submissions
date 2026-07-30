func rotate(nums []int, k int) {
	if len(nums) == k || k == 0 {
		return
	}

	if len(nums) % k != 0 {
		move(nums, k, 0)
		return
	}

	for i := 0; i < k; i++ {
		move(nums, k, i)
	}
}

func move(nums []int, k, i int) {
	startedI := i
	mem := nums[i]
	for {
		j := (i + k)%len(nums)
		mem, nums[j] = nums[j], mem
		i = j

		if startedI == i {
			break
		}
	}
}