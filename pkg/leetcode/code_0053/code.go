package code_0053

// @1 @2 @3 @4 @5 @6 @7 @8
// result = max(max(...@1), max(...@2), max(...@3), ...)
// max(...@3) = @3 || @3 + max(...@2)
// max(...@1) = nums[0]
// max(...@2) = nums[1] || nums[1] + max(...@1)
// ......

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	result := dp[0]
	for i:=1;i<len(nums);i++ {
		dp[i] = max(nums[i], dp[i-1] + nums[i])
		result = max(result, dp[i])
	}
	return result
}
func max(a int, b int) int  {
	if a > b {
		return a
	}
	return b
}