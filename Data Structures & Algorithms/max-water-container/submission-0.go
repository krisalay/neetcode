func maxArea(heights []int) int {
	result := 0
	left, right := 0, len(heights) - 1
	for left < right {
		height := min(heights[left], heights[right])
		area := height * (right - left)
		result = max(result, area)
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}