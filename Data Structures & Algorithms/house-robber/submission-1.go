func rob(nums []int) int {
	n := len(nums)
	dp1, dp2 := 0, nums[0]
	for i := 2; i <= n; i++ {
		dp1, dp2 = dp2, max(dp1+nums[i-1], dp2)
	}
	return dp2
}
