func twoSum(nums []int, target int) []int {
    diff := make(map[int]int)
    for i, num := range nums {
        diff[target-num] = i
    }

    for i, num := range nums {
        if idx, ok := diff[num]; ok && i != idx {
            return []int{i, idx}
        }
    }
    return nil
}
