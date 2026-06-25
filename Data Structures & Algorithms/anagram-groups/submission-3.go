import "slices"

func groupAnagrams(strs []string) [][]string {
    hashMap := make(map[string][]string)

	for _, str := range strs {
		strRune := []rune(str)
		slices.Sort(strRune)
		sortedStr := string(strRune)

		hashMap[sortedStr] = append(hashMap[sortedStr], str)
	}

	result := make([][]string, 0, len(hashMap))
	for _, bucket := range hashMap {
		result = append(result, bucket)
	}

	return result
}


// {
//     "act": ["act", "cat"],
//     "opst": ["pots", "tops", "stop"],
//     "aht": ["hat"]
// }

// TC = O(n) * O(mlogm)
// SC = O(n)