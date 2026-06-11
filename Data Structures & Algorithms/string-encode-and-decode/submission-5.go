type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteString("#")
		sb.WriteString(str)
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	ptr := 0
	var sizeBuilder strings.Builder
	result := []string{}
	for ptr < len(encoded) {
		for encoded[ptr] != '#' {
			sizeBuilder.WriteString(string(encoded[ptr]))
			ptr++
		}
		ptr++
		size, _ := strconv.Atoi(sizeBuilder.String())
		result = append(result, encoded[ptr:(ptr+size)])
		ptr = ptr + size
		sizeBuilder.Reset()
	}
	return result
}
