func groupAnagrams(strs []string) [][]string {
    groups := make(map[[26]int][]string)
    for _, str := range strs {
        var key [26]int
        for _, char := range str {
            key[char-'a']++
        }
        groups[key] = append(groups[key], str)
    }

    res := make([][]string, 0, len(groups))
    for _, group := range groups {
        res = append(res, group)
    }
    return res
}
