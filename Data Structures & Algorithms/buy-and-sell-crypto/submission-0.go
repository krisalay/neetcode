func maxProfit(prices []int) int {
	maxSellingPrice := prices[len(prices)-1]
	maxProfit := math.MinInt
	for i := len(prices)-1; i >= 0; i-- {
		maxProfit = max(maxProfit, maxSellingPrice - prices[i])
		maxSellingPrice = max(maxSellingPrice, prices[i])
	}
	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}