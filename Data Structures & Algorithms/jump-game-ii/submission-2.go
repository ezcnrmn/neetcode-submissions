func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	for i := n-2; i >= 0; i-- {
		jumps := n
		for j := i+1; j < n && j <= i + nums[i]; j++ {
			jumps = min(jumps, dp[j])
		}
		dp[i] = 1 + jumps
	}

	return dp[0]
}
