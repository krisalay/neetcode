func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	ans := nums[0]
	for left <= right {
		mid := (left+right)/2
		if nums[left] <= nums[mid] { // left is sorted
			ans = min(ans, nums[left])
			left = mid + 1
		} else { // right is sorted
			ans = min(ans, nums[mid])
			right = mid - 1
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
