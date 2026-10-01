func longestCommonSubsequence(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	selected1 := make([]bool, n)
	selected2 := make([]bool, m)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			var sameLetter int
			if text1[i-1] == text2[j-1] && !selected1[i-1] && !selected2[j-1] {
				sameLetter = 1
				selected1[i-1] = true
				selected2[j-1] = true
			}
			dp[i][j] = max(dp[i][j-1], dp[i-1][j-1], dp[i-1][j]) + sameLetter
		}
	}

	return dp[n][m]
}
