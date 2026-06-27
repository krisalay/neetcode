func findMin(nums []int) int {
    left, right := 0, len(nums)-1
    res := nums[0]
    
    for left <= right {
        mid := left + (right-left)/2
        res = min(res, nums[mid])
        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
