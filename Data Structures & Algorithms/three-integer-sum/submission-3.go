func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    var result [][]int

    for i, num := range nums {
        if num > 0 {
            break
        }
        if i > 0 && num == nums[i-1] {
            continue
        }

        l, r := i+1, len(nums)-1
        for l < r {
            sum := num + nums[l] + nums[r]
            if sum == 0 {
                result = append(result, []int{num, nums[l], nums[r]})
                l++
                r--
                for l < r && nums[l] == nums[l-1] {
                    l++
                }
            } else if sum < 0 {
                l++
            } else {
                r--
            }
        }
    }
    return result
}