package code_122

// 有得赚，就卖
func maxProfit(prices []int) int {
	maxValue := 0
	for i := 0; i < len(prices)-1; i++ {
		t := prices[i+1] - prices[i]
		if t > 0 {
			maxValue = maxValue + t
		}
	}
	return maxValue
}
