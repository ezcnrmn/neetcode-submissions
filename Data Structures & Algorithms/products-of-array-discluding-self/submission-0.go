func productExceptSelf(nums []int) []int {
	var (
		befores = make([]int, len(nums))
		afters = make([]int, len(nums))
		result = make([]int, len(nums))
	)
	befores[0]=1
	afters[len(nums)-1]=1

	for i := 1; i < len(nums); i++ {
		befores[i] = befores[i-1]*nums[i-1]
	}
	for i := len(nums)-2; i >= 0; i-- {
		afters[i] = afters[i+1]*nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		result[i] = afters[i]*befores[i]
	}

	return result
}
