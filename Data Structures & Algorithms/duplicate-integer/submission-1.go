func hasDuplicate(nums []int) bool {
    uniqueMap := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := uniqueMap[num]; !ok {
			uniqueMap[num] = struct{}{}
			continue
		}
		return true
	}
	return false
}
