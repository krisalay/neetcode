func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freqMap := make(map[rune]int)
	for _, char := range s {
		freqMap[char]++
	}

	for _, char := range t {
		freqMap[char]--
	}

	for _, freq := range freqMap {
		if freq != 0 {
			return false
		}
	}

	return true
}
