func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		if s[n-i] == '0' { continue }

		num := string([]byte{s[n-i], s[n-i+1]})
		parsed, err := strconv.Atoi(num)
		if err == nil && parsed <= 26 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[n]
}
