func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	n := len(nums)
	dp := make([]int, n)
	dp[1] = nums[0]
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2] + nums[i-1])
	}
	res1 := dp[n-1]

	dp[1] = nums[1]
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2] + nums[i])
	}

	return max(res1, dp[n-1])
}
