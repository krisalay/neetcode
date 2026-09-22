type Solution struct{}

const delimiter = '§'

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
		sb.WriteString(string(delimiter))
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	var res []string
	var sb strings.Builder
	for _, char := range encoded {
		if char == delimiter {
			res = append(res, sb.String())
			sb.Reset()
		} else {
			sb.WriteString(string(char))
		}
	}
	return res
}
