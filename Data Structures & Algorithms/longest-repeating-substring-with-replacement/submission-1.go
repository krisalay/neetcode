func characterReplacement(s string, k int) int {
    maxLength := 0
    left, right := 0, 0
    freqMap := make(map[byte]int)

    for right < len(s) {
        freqMap[s[right]]++
        maxFreq := getMostFrequentCount(freqMap)
        windowLength := right-left+1
        if windowLength - maxFreq <= k {
            maxLength = max(maxLength, windowLength)
        } else {
            freqMap[s[left]]--
            left++
        }
        right++
    }

    return maxLength
}

func getMostFrequentCount(freqMap map[byte]int) int {
    maxFreq := 0
    for _, count := range freqMap {
        maxFreq = max(maxFreq, count)
    }
    return maxFreq
}

func max(a, b int) int {
    if a > b {
        return a    
    }
    return b
}