func lengthOfLongestSubstring(s string) int {
	idxMap := make(map[byte]int)
	left := 0
	longest := 0
	for right := 0; right < len(s); right++ {
		currByte := s[right]
		if idx, ok := idxMap[currByte]; ok && idx >= left {
			left = idx+1
		}
		idxMap[currByte] = right
		window := right - left + 1
		if window > longest {
			longest = window
		}
	}
	return longest
}
