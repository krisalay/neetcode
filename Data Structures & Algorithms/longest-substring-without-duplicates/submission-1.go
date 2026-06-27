func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	left, right := 0, 1
	sRune := []rune(s)
	placeholderMap := make(map[rune]int)
	placeholderMap[sRune[0]] = 0
	maxLength := 1
	for right < len(s) {
		idx, ok := placeholderMap[sRune[right]]
		if ok && idx >= left {
			maxLength = max(maxLength, right-left)
			left = idx + 1
		}
		placeholderMap[sRune[right]] = right
		right++
	}
	return max(maxLength, right-left)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}