func hasDuplicate(nums []int) bool {
	uniqueMap := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := uniqueMap[num]; ok {
			return true
		}
		uniqueMap[num] = struct{}{}
	}
	return false
}