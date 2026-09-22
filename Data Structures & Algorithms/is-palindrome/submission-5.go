func isPalindrome(str string) bool {
	s := strings.ToLower(str)
	left, right := 0, len(s)-1
	for left < right {
		for left < len(str) && !isAlphanumeric(rune(s[left])) {
			left++
		}
		for right >= 0 && !isAlphanumeric(rune(s[right])) {
			right--
		}
		if left >= right {
            break
        }
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphanumeric(char rune) bool {
	return unicode.IsDigit(char) || unicode.IsLetter(char)
}