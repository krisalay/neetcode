func combinationSum(nums []int, target int) [][]int {
    var result [][]int

	var dfs func(idx int, curr []int, total int)
	dfs = func(idx int, curr []int, total int) {
		if total == target {
			temp := make([]int, len(curr))
			copy(temp, curr)
			result = append(result, temp)
			return
		}

		if idx >= len(nums) || total > target {
			return
		}

		curr = append(curr, nums[idx])
		dfs(idx, curr, total + nums[idx])

		curr = curr[:len(curr)-1]
		dfs(idx+1, curr, total)
	}
	dfs(0, []int{}, 0)
	return result
}
