func maxProfit(prices []int) int {
	currMax := prices[len(prices)-1]
	var result int
	for i := len(prices)-1; i >= 0; i-- {
		profit := currMax - prices[i]
		if profit > result {
			result = profit
		}
		if prices[i] > currMax {
			currMax = prices[i]
		}
	}
	return result
}
