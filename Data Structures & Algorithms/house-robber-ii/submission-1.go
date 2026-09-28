func rob(nums []int) int {
	n := len(nums)
	dp1, dp2 := 0, nums[0]
	for i := 2; i < n; i++ {
		dp1, dp2 = dp2, max(dp1+nums[i-1], dp2)
	}
	solution1 := dp2
	dp1, dp2 = 0, nums[1]
	for i := 3; i <= n; i++ {
		dp1, dp2 = dp2, max(dp1+nums[i-1], dp2)
	}
	return max(solution1, dp2)
}
