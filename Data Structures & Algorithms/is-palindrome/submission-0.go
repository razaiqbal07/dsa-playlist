func isPalindrome(s string) bool {
	var b strings.Builder

	for _, char := range s {
		if unicode.IsLetter(rune(char)) || unicode.IsDigit(rune(char)) {
			b.WriteRune(rune(char))
		}
	}

	s = strings.ToLower(b.String())
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--

	}
	return true
}
