type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
		sb.WriteString("é")
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}
	var sb strings.Builder
	for _, char := range encoded {
		if char != 'é' {
			sb.WriteString(string(char))
		} else {
			result = append(result, sb.String())
			sb.Reset()
		}
	}
	return result
}
