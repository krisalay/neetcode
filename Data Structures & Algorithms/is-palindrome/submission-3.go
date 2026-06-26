func isPalindrome(s string) bool {
	str := strings.ToLower(s)
	left, right := 0, len(s)-1

	for left < right {
		for isIndexValid(right, len(s)) && !isAlphaNumeric(rune(str[right])) {
			right--
		}
		for isIndexValid(left, len(s)) && !isAlphaNumeric(rune(str[left])) {
			left++
		}
		if isIndexValid(left, len(s)) && isIndexValid(right, len(s)) && str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isIndexValid(idx, arrLen int) bool {
	return idx >= 0 && idx < arrLen
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLetter(r)
}