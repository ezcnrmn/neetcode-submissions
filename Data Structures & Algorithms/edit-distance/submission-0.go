func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	dp := make([][]int, n+1)
	for i := range n+1 {
		dp[i] = make([]int, m+1)
		dp[i][0] = i
	}
	for i := range m+1 {
		dp[0][i] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			} 

			dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
		}
	}

	return dp[n-1][m-1]
}
