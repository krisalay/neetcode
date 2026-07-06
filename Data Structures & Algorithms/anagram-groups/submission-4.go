func groupAnagrams(strs []string) [][]string {
	buckets := make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int
		for _, char := range str {
			count[char-'a']++
		}
		buckets[count] = append(buckets[count], str)
	}
	var result [][]string
	for _, bucket := range buckets {
		result = append(result, bucket)
	}
	return result
}
