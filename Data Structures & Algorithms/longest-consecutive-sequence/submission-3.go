func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	output := 0
	for _, num := range nums {
		if isBeginningOfSequence(num, set) {
			output = max(output, sequenceLength(num, set))
		}
	}
	return output
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isBeginningOfSequence(num int, set map[int]struct{}) bool {
	if _, ok := set[num-1]; !ok {
		return true
	}
	return false
}

func sequenceLength(num int, set map[int]struct{}) int {
	output := 0
	for true {
		if _, ok := set[num]; ok {
			output++
			num++
		} else {
			break
		}
	}
	return output
}