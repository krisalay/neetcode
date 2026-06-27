func twoSum(nums []int, target int) []int {
	checker := make(map[int]int)

	for idx , num := range nums {
		diff := target -num

		i, ok := checker[diff]
		if ok {
			return []int{i,idx}
		}

		checker[num] = idx
	}

	return []int{0,0}
} 