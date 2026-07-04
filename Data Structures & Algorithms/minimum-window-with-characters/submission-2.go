
func minWindow(s string, t string) string {
	temp := t
	left := 0
	right := 0
	minSubstring := ""
	for idx, char := range s {
		if strings.Contains(temp, string(char)) {
			temp = strings.Replace(temp, string(char), "", 1)
			left = idx
			right = idx + 1
			if temp == "" {
				if minSubstring != "" {
					current := ""
					if right > len(s) {
						current = s[left:]
					} else {
						current = s[left:right]
					}
					if len(current) < len(minSubstring) {
						minSubstring = current
					}
				} else {
					if right > len(s) {
						minSubstring = s[left:]
					} else {
						minSubstring = s[left:right]
					}

				}
				break
			}

			for right < len(s) {
				if strings.Contains(temp, string(s[right])) {
					temp = strings.Replace(temp, string(s[right]), "", 1)
				}
				if temp == "" {
					if minSubstring != "" {
						current := s[left : right+1]
						if len(current) < len(minSubstring) {
							minSubstring = current
						}
					} else {
						minSubstring = s[left : right+1]

					}
					break
				}
				right++
			}
		}
		temp = t
		right = 0
	}
	return minSubstring
}