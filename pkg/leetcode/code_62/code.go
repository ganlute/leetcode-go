package code_62

// dp[i][j] = dp[i-1][j] + dp[i][j-1]
// dp[0][j] = 1
// dp[i][0] = 1
func uniquePaths(m int, n int) int {
	dp := make([][]int, m, m)
	for i := range dp {
		dp[i] = make([]int, n, n)
	}
	// 初始化
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	// dp
	for j := 1; j < n; j++ {
		for i := 1; i < m; i++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
