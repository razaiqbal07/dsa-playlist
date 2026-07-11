// Change the name from Codec to Solution
type Solution struct{}

// Encodes a list of strings to a single string.
func (c *Solution) Encode(strs []string) string {
	var sb strings.Builder

	for _, s := range strs {
		sb.WriteString(strconv.Itoa(len(s)))
		sb.WriteByte('#')
		sb.WriteString(s)
	}

	return sb.String()
}

// Decodes a single string to a list of strings.
func (c *Solution) Decode(s string) []string {
	var res []string
	i := 0

	for i < len(s) {
		j := i
		for j < len(s) && s[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(s[i:j])
		j++ // Move past the '#'

		res = append(res, s[j:j+length])
		
		i = j + length
	}

	return res
}