package code_121

func maxProfit(prices []int) int {
	maxValue := 0
	i:=0
	j := len(prices)-1
	for i < j {
		maxValue = max(maxValue, prices[j] - prices[i])
		if prices[i] > prices[i+1] {
			i = i +1
			continue
		}
		if prices[j] < prices[j-1] {
			j = j -1
			continue
		}
		
	}
	return maxValue
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
