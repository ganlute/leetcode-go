package code_0121

// 前i天的利润 = max (前i-1天的利润, 第i天的价格 - 前i-1天最小价格)

func maxProfit(prices []int) int {
	minPrice := make([]int, 0)
	maxPro := make([]int, 0)
	// 初始化
	minPrice = append(minPrice, prices[0])
	maxPro = append(maxPro, 0)

	for i := 1; i < len(prices); i++ {
		minPrice = append(minPrice, min(minPrice[i-1], prices[i]))
		maxPro = append(maxPro, max(maxPro[i-1], prices[i]-minPrice[i-1]))
	}
	return maxPro[len(prices)-1]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
