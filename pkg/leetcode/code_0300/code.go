package code_0300

// dp[i] = max( dp[k] + 1 if nums[k] < nums[i]) (k = 0 ~ i-1)
// dp [0] = 1

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for k := 0; k <= i-1; k++ {
			if nums[k] < nums[i] {
				dp[i] = max(dp[i], dp[k]+1)
			}
		}
	}
	maxV := dp[0]
	for i := 1; i < len(nums); i++ {
		maxV = max(dp[i], maxV)
	}
	return maxV
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
