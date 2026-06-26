func trap(height []int) int {
	maxL := make([]int, len(height))
	maxR := make([]int, len(height))

	tempMaxL := 0
	for i, num := range height {
		maxL[i] = tempMaxL
		tempMaxL = max(tempMaxL, num)
	}
	tempMaxR := 0
	for i := len(height)-1; i >= 0; i-- {
		maxR[i] = tempMaxR
		tempMaxR = max(tempMaxR, height[i])
	}

	result := 0
	for i, h := range height {
		trapped := min(maxL[i], maxR[i]) - h
		if trapped > 0 {
			result += trapped
		}
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}