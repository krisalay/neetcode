type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
		sb.WriteString(".")
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}
	var sb strings.Builder
	for _, char := range encoded {
		if char == '.' {
			result = append(result, sb.String())
			sb.Reset()
		} else {
			sb.WriteString(string(char))
		}
	}
	return result
}
