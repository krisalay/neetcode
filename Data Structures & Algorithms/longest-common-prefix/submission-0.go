

func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)

	res := []byte{}
	for i, j := 0, 0; i < len(strs[0]) && j < len(strs[len(strs)-1]);i,j=i+1,j+1 {
		if strs[0][i] != strs[len(strs)-1][j] {
			return string(res)
		}
		res = append(res, strs[0][i])
	}
	return string(res)
}
