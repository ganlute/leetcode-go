package code_0322

// dp[j] = min( dp[j-coin]+1 )

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}
	// 初始化
	dp[0] = 0
	for i := range coins {
		if coins[i] > amount {
			continue
		}
		dp[coins[i]] = 1
	}
	for i := 1; i <= amount; i++ {
		if dp[i] != -1 {
			continue
		}
		minV := amount + 1
		for j := range coins {
			if i-coins[j] < 0 {
				continue
			}
			if dp[i-coins[j]] == -1 {
				continue
			}
			minV = min(dp[i-coins[j]]+1, minV)
		}
		if minV == amount+1 {
			dp[i] = -1
		} else {
			dp[i] = minV
		}
	}
	return dp[amount]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
