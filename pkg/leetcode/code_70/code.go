package code_70

// dp[n] = dp[n-1] + dp[n-2]
// dp[1] = 1
// dp[2] = 2
// dp[0] = 1 // 虚拟
func climbStairs(n int) int {
	dp := make([]int, 0)
	dp = append(dp, 1)
	dp = append(dp, 1)
	for i := 2; i <= n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}
