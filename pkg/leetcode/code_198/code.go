package code_198

// 方法一
// 对一个屋子，小偷可以偷 或者 不偷
// 如果偷，则对第3个屋子往后递归
// 如果不偷，则对第2哥屋子往后递归
// 方法一复杂度超了
//func rob(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	if len(nums) == 1 {
//		return nums[0]
//	}
//	if len(nums) == 2 {
//		return max(nums[0], nums[1])
//	}
//	return max(nums[0]+rob(nums[2:]), rob(nums[1:]))
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

// 方法二
// dp[i] = max (dp[i-2] + nums[i] , dp[i-1] )
// dp[0] = nums[0]
// dp[1] = max(nums[0], nums[1])

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums), len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
