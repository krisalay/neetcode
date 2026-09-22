func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	longest := 0
	for num, _ := range set {
		count := 0
		if _, ok := set[num-1]; !ok {
			for set[num] {
				count++
				delete(set, num)
				num++
			}
			if count > longest {
				longest = count
			}
		}
	}
	return longest
}
