

func checkInclusion(s1 string, s2 string) bool {
	temp := s1
	for idx, r := range s2 {
		if strings.Contains(temp, string(r)) {
			temp = strings.Replace(temp, string(r), "", 1)
			if temp == "" {
				return true
			}
			substring := s2[idx+1:]
			for _, ch := range substring {
				if strings.Contains(temp, string(ch)) {
					temp = strings.Replace(temp, string(ch), "", 1)
					if temp == "" {
						return true
					}
				} else {
					break
				}
			}
		}
		temp = s1
	}

	return len(temp) == 0
}