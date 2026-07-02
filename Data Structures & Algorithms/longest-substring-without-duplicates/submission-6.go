

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	right := 0
	str := ""
	for right < len(s) {
		if strings.Contains(str, string(s[right])) {
			i := strings.IndexByte(str, s[right])
			str = str[i+1:]
		}
		str += string(s[right])
		maxLen = max(maxLen, len(str))
		right++
	}
	return maxLen
}