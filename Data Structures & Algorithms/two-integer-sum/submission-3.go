func twoSum(nums []int, target int) []int {
    hashmap := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		idx, ok := hashmap[diff]
		if !ok {
			hashmap[num] = i
		} else if idx != i {
			return []int{idx, i}
		}
	}
	return []int{}
}
