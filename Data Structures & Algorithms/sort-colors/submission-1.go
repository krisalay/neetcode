func sortColors(nums []int) {
    left, right := 0, len(nums)-1
	read := 0

	for read <= right {
		if nums[read] == 2 {
			nums[right], nums[read] = nums[read], nums[right]
			right--
		} else if nums[read] == 0 {
			nums[left], nums[read] = nums[read], nums[left]
			left++
			read++
		} else {
			read++
		}
	}
}
