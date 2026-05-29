func groupAnagrams(strs []string) [][]string {
    hashMap := make(map[string][]string)

    for _, str := range strs {
        key := sortStrings(str)
        hashMap[key] = append(hashMap[key], str)
    }

    result := [][]string{}
    for _, v := range hashMap {
        result = append(result, v)
    }

    return result
}

func sortStrings(s string) string {
    chars := []byte(s)

    sort.Slice(chars, func(i, j int) bool {
        return chars[i] < chars[j]
    })
    return string(chars)
}
