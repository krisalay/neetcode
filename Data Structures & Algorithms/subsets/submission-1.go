func subsets(nums []int) [][]int {
	res := [][]int{}

	var subset []int
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx >= len(nums) {
			copy := append([]int{}, subset...)
			res = append(res, copy)
			return
		}

		subset = append(subset, nums[idx])
		dfs(idx+1)

		subset = subset[:len(subset)-1]
		dfs(idx+1)
	}
	dfs(0)
	return res
}
