func longestConsecutive(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	sort.Ints(nums)
	globalMax, currMax := 0, 1
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff == 1 {
			currMax++
		} else if diff > 1 {
			globalMax = max(globalMax, currMax)
			currMax = 1
		}
	}

	return max(globalMax, currMax)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
