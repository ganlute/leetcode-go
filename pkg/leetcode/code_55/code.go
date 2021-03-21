package code_55

// dp[i] = (dp[i-1] && nums[i-1] >= 1) || ... || (dp[i-j] && nums[i-j] >= j)
// dp[0] = true
func canJump(nums []int) bool {
	dp := make([]bool, len(nums), len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		dp[i] = false
		for j := 1; i-j >= 0; j++ {
			dp[i] = dp[i-1] && (nums[i-j] >= j) || dp[i]
			if dp[i] {
				break
			}
		}
	}
	return dp[len(nums)-1]
}
